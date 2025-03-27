package email

import (
	"bytes"
	"context"
	"html/template"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	texttemplate "text/template"

	"github.com/juicycleff/frank/config"
	"github.com/juicycleff/frank/ent"
	"github.com/juicycleff/frank/pkg/errors"
	"github.com/juicycleff/frank/pkg/logging"
)

// TemplateManager manages email templates
type TemplateManager struct {
	config        *config.Config
	logger        logging.Logger
	htmlTemplates map[string]*template.Template
	textTemplates map[string]*texttemplate.Template
	repo          TemplateRepository
	mu            sync.RWMutex // For thread safety when accessing templates
}

// TemplateRepository provides access to email template storage
type TemplateRepository interface {
	// Create creates a new email template
	Create(ctx context.Context, input TemplateRepositoryCreateInput) (*ent.EmailTemplate, error)

	// GetByID retrieves an email template by ID
	GetByID(ctx context.Context, id string) (*ent.EmailTemplate, error)

	// GetByTypeAndOrganization retrieves an email template by type and organization
	GetByTypeAndOrganization(ctx context.Context, templateType, organizationID, locale string) (*ent.EmailTemplate, error)

	// List retrieves email templates with pagination
	List(ctx context.Context, input TemplateRepositoryListInput) ([]*ent.EmailTemplate, int, error)

	// Update updates an email template
	Update(ctx context.Context, id string, input TemplateRepositoryUpdateInput) (*ent.EmailTemplate, error)

	// Delete deletes an email template
	Delete(ctx context.Context, id string) error
}

// TemplateRepositoryCreateInput represents input for creating an email template
type TemplateRepositoryCreateInput struct {
	Name           string
	Subject        string
	Type           string
	HTMLContent    string
	TextContent    string
	OrganizationID string
	Active         bool
	System         bool
	Locale         string
	Metadata       map[string]interface{}
}

// TemplateRepositoryUpdateInput represents input for updating an email template
type TemplateRepositoryUpdateInput struct {
	Name        *string
	Subject     *string
	HTMLContent *string
	TextContent *string
	Active      *bool
	Locale      *string
	Metadata    map[string]interface{}
}

// TemplateRepositoryListInput represents input for listing email templates
type TemplateRepositoryListInput struct {
	Offset         int
	Limit          int
	Type           string
	OrganizationID string
	Locale         string
}

// NewTemplateManager creates a new template manager
func NewTemplateManager(repo TemplateRepository, cfg *config.Config, logger logging.Logger) *TemplateManager {
	manager := &TemplateManager{
		config:        cfg,
		logger:        logger,
		htmlTemplates: make(map[string]*template.Template),
		textTemplates: make(map[string]*texttemplate.Template),
		repo:          repo,
	}

	// Load templates
	if err := manager.LoadTemplates(); err != nil {
		logger.Error("Failed to load email templates", logging.Error(err))
	}

	return manager
}

type templatePersist struct {
	name        string
	subject     string
	htmlContent string
	textContent string
	locale      string
}

// LoadTemplates loads templates from the templates directory
func (m *TemplateManager) LoadTemplates() error {
	// Get template paths
	templatePath := m.config.Templates.EmailPath
	if templatePath == "" {
		templatePath = "./web/templates/email"
	}

	// Check if directory exists
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		m.logger.Warn("Email templates directory not found", logging.String("path", templatePath))
		return nil
	}

	// Lock for writing
	m.mu.Lock()
	defer m.mu.Unlock()

	// Clear existing templates
	m.htmlTemplates = make(map[string]*template.Template)
	m.textTemplates = make(map[string]*texttemplate.Template)

	// Keep track of templates to persist to database
	templatesToPersist := make(map[string]templatePersist)

	// Walk template directory
	err := filepath.Walk(templatePath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Get file extension
		ext := strings.ToLower(filepath.Ext(path))

		// Only process .html and .txt files
		if ext != ".html" && ext != ".txt" {
			return nil
		}

		// Read file
		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		// Get template name (file name without extension)
		baseName := strings.TrimSuffix(filepath.Base(path), ext)

		// Check if this is a template in a locale directory
		relPath, err := filepath.Rel(templatePath, path)
		if err != nil {
			return err
		}

		// Extract template name and locale
		parts := strings.Split(relPath, string(filepath.Separator))
		locale := "en" // Default locale
		templateName := baseName

		if len(parts) > 1 {
			// Format: locale/templateName.ext
			locale = parts[0]
			templateName = strings.TrimSuffix(parts[1], ext)
		}

		// Parse template
		if ext == ".html" {
			tmpl, err := template.New(templateName).Parse(string(content))
			if err != nil {
				m.logger.Warn("Failed to parse HTML template",
					logging.String("name", templateName),
					logging.Error(err),
				)
				return nil
			}
			m.htmlTemplates[templateName] = tmpl

			// Store for database persistence
			if t, exists := templatesToPersist[templateName]; exists {
				t.htmlContent = string(content)
				templatesToPersist[templateName] = t
			} else {
				templatesToPersist[templateName] = struct {
					name        string
					subject     string
					htmlContent string
					textContent string
					locale      string
				}{
					name:        templateName,
					subject:     "[" + strings.Title(templateName) + "]", // Default subject
					htmlContent: string(content),
					locale:      locale,
				}
			}
		} else {
			tmpl, err := texttemplate.New(templateName).Parse(string(content))
			if err != nil {
				m.logger.Warn("Failed to parse text template",
					logging.String("name", templateName),
					logging.Error(err),
				)
				return nil
			}
			m.textTemplates[templateName] = tmpl

			// Store for database persistence
			if t, exists := templatesToPersist[templateName]; exists {
				t.textContent = string(content)
				templatesToPersist[templateName] = t
			} else {
				templatesToPersist[templateName] = struct {
					name        string
					subject     string
					htmlContent string
					textContent string
					locale      string
				}{
					name:        templateName,
					subject:     "[" + strings.Title(templateName) + "]", // Default subject
					textContent: string(content),
					locale:      locale,
				}
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	m.logger.Info("Loaded email templates",
		logging.Int("html_count", len(m.htmlTemplates)),
		logging.Int("text_count", len(m.textTemplates)),
	)

	return m.persistTemplates(templatesToPersist)
}

// RenderHTML renders an HTML template with data
func (m *TemplateManager) persistTemplates(templatesToPersist map[string]templatePersist) error {
	// Persist templates to database if repository is available
	if repo, ok := interface{}(m).(interface{ GetRepository() TemplateRepository }); ok {
		templateRepo := repo.GetRepository()

		// Create a background context for database operations
		ctx := context.Background()

		// Persist each template
		for templateName, template := range templatesToPersist {
			// Skip templates with missing HTML content
			if template.htmlContent == "" {
				continue
			}

			// Try to find existing template in database
			existingTemplate, err := templateRepo.GetByTypeAndOrganization(ctx, templateName, "", template.locale)

			if err != nil {
				if errors.IsNotFound(err) {
					// Create new template
					_, err = templateRepo.Create(ctx, TemplateRepositoryCreateInput{
						Name:        template.name,
						Subject:     template.subject,
						Type:        templateName,
						HTMLContent: template.htmlContent,
						TextContent: template.textContent,
						Locale:      template.locale,
						Active:      true,
						System:      true, // Mark as system template
					})

					if err != nil {
						m.logger.Error("Failed to persist template to database",
							logging.String("name", templateName),
							logging.Error(err),
						)
					} else {
						m.logger.Info("Created new template in database",
							logging.String("name", templateName),
							logging.String("locale", template.locale),
						)
					}
				} else {
					m.logger.Error("Failed to check for existing template",
						logging.String("name", templateName),
						logging.Error(err),
					)
				}
				continue
			}

			// Update existing template if content has changed
			if existingTemplate.HTMLContent != template.htmlContent ||
				existingTemplate.TextContent != template.textContent {

				htmlContent := template.htmlContent
				textContent := template.textContent

				_, err = templateRepo.Update(ctx, existingTemplate.ID, TemplateRepositoryUpdateInput{
					HTMLContent: &htmlContent,
					TextContent: &textContent,
				})

				if err != nil {
					m.logger.Error("Failed to update template in database",
						logging.String("name", templateName),
						logging.Error(err),
					)
				} else {
					m.logger.Info("Updated template in database",
						logging.String("name", templateName),
						logging.String("locale", template.locale),
					)
				}
			}
		}
	}

	return nil
}

// RenderHTML renders an HTML template with data
func (m *TemplateManager) RenderHTML(content string, data map[string]interface{}) (string, error) {
	// Parse template
	tmpl, err := template.New("dynamic").Parse(content)
	if err != nil {
		return "", err
	}

	// Render template
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// RenderText renders a text template with data
func (m *TemplateManager) RenderText(content string, data map[string]interface{}) (string, error) {
	// Parse template
	tmpl, err := texttemplate.New("dynamic").Parse(content)
	if err != nil {
		return "", err
	}

	// Render template
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
