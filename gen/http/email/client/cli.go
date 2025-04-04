// Code generated by goa v3.20.0, DO NOT EDIT.
//
// email HTTP client CLI support package
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	email "github.com/juicycleff/frank/gen/email"
	goa "goa.design/goa/v3/pkg"
)

// BuildListTemplatesPayload builds the payload for the email list_templates
// endpoint from CLI flags.
func BuildListTemplatesPayload(emailListTemplatesOffset string, emailListTemplatesLimit string, emailListTemplatesType string, emailListTemplatesOrganizationID string, emailListTemplatesLocale string, emailListTemplatesJWT string) (*email.ListTemplatesPayload, error) {
	var err error
	var offset int
	{
		if emailListTemplatesOffset != "" {
			var v int64
			v, err = strconv.ParseInt(emailListTemplatesOffset, 10, strconv.IntSize)
			offset = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for offset, must be INT")
			}
			if offset < 0 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("offset", offset, 0, true))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var limit int
	{
		if emailListTemplatesLimit != "" {
			var v int64
			v, err = strconv.ParseInt(emailListTemplatesLimit, 10, strconv.IntSize)
			limit = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for limit, must be INT")
			}
			if limit < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 1, true))
			}
			if limit > 100 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 100, false))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var type_ *string
	{
		if emailListTemplatesType != "" {
			type_ = &emailListTemplatesType
		}
	}
	var organizationID *string
	{
		if emailListTemplatesOrganizationID != "" {
			organizationID = &emailListTemplatesOrganizationID
		}
	}
	var locale *string
	{
		if emailListTemplatesLocale != "" {
			locale = &emailListTemplatesLocale
		}
	}
	var jwt *string
	{
		if emailListTemplatesJWT != "" {
			jwt = &emailListTemplatesJWT
		}
	}
	v := &email.ListTemplatesPayload{}
	v.Offset = offset
	v.Limit = limit
	v.Type = type_
	v.OrganizationID = organizationID
	v.Locale = locale
	v.JWT = jwt

	return v, nil
}

// BuildCreateTemplatePayload builds the payload for the email create_template
// endpoint from CLI flags.
func BuildCreateTemplatePayload(emailCreateTemplateBody string, emailCreateTemplateJWT string) (*email.CreateTemplatePayload, error) {
	var err error
	var body CreateTemplateRequestBody
	{
		err = json.Unmarshal([]byte(emailCreateTemplateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"active\": false,\n      \"html_content\": \"\\u003chtml\\u003e\\u003cbody\\u003e\\u003ch1\\u003eWelcome!\\u003c/h1\\u003e\\u003cp\\u003eHello {{name}}\\u003c/p\\u003e\\u003c/body\\u003e\\u003c/html\\u003e\",\n      \"locale\": \"Aut et numquam qui.\",\n      \"metadata\": {\n         \"Autem deleniti.\": \"Laudantium maiores debitis nam saepe.\",\n         \"Ea amet voluptatem.\": \"Voluptatem nobis iste occaecati aut eos voluptas.\",\n         \"Est voluptatibus qui.\": \"Amet rerum ut.\"\n      },\n      \"name\": \"Welcome Email\",\n      \"organization_id\": \"Quam explicabo saepe ipsam ea porro.\",\n      \"subject\": \"Welcome to our platform\",\n      \"system\": false,\n      \"text_content\": \"Welcome! Hello {{name}}\",\n      \"type\": \"welcome\"\n   }'")
		}
	}
	var jwt *string
	{
		if emailCreateTemplateJWT != "" {
			jwt = &emailCreateTemplateJWT
		}
	}
	v := &email.CreateTemplatePayload{
		Name:           body.Name,
		Subject:        body.Subject,
		Type:           body.Type,
		HTMLContent:    body.HTMLContent,
		TextContent:    body.TextContent,
		OrganizationID: body.OrganizationID,
		Active:         body.Active,
		System:         body.System,
		Locale:         body.Locale,
	}
	{
		var zero bool
		if v.Active == zero {
			v.Active = true
		}
	}
	{
		var zero bool
		if v.System == zero {
			v.System = false
		}
	}
	{
		var zero string
		if v.Locale == zero {
			v.Locale = "en"
		}
	}
	if body.Metadata != nil {
		v.Metadata = make(map[string]any, len(body.Metadata))
		for key, val := range body.Metadata {
			tk := key
			tv := val
			v.Metadata[tk] = tv
		}
	}
	v.JWT = jwt

	return v, nil
}

// BuildGetTemplatePayload builds the payload for the email get_template
// endpoint from CLI flags.
func BuildGetTemplatePayload(emailGetTemplateID string, emailGetTemplateJWT string) (*email.GetTemplatePayload, error) {
	var id string
	{
		id = emailGetTemplateID
	}
	var jwt *string
	{
		if emailGetTemplateJWT != "" {
			jwt = &emailGetTemplateJWT
		}
	}
	v := &email.GetTemplatePayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildGetTemplateByTypePayload builds the payload for the email
// get_template_by_type endpoint from CLI flags.
func BuildGetTemplateByTypePayload(emailGetTemplateByTypeType string, emailGetTemplateByTypeOrganizationID string, emailGetTemplateByTypeLocale string, emailGetTemplateByTypeJWT string) (*email.GetTemplateByTypePayload, error) {
	var type_ string
	{
		type_ = emailGetTemplateByTypeType
	}
	var organizationID *string
	{
		if emailGetTemplateByTypeOrganizationID != "" {
			organizationID = &emailGetTemplateByTypeOrganizationID
		}
	}
	var locale string
	{
		if emailGetTemplateByTypeLocale != "" {
			locale = emailGetTemplateByTypeLocale
		}
	}
	var jwt *string
	{
		if emailGetTemplateByTypeJWT != "" {
			jwt = &emailGetTemplateByTypeJWT
		}
	}
	v := &email.GetTemplateByTypePayload{}
	v.Type = type_
	v.OrganizationID = organizationID
	v.Locale = locale
	v.JWT = jwt

	return v, nil
}

// BuildUpdateTemplatePayload builds the payload for the email update_template
// endpoint from CLI flags.
func BuildUpdateTemplatePayload(emailUpdateTemplateBody string, emailUpdateTemplateID string, emailUpdateTemplateJWT string) (*email.UpdateTemplatePayload, error) {
	var err error
	var body UpdateTemplateRequestBody
	{
		err = json.Unmarshal([]byte(emailUpdateTemplateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"template\": {\n         \"active\": false,\n         \"html_content\": \"Qui blanditiis provident libero.\",\n         \"locale\": \"Ut dolorum ipsa eos neque.\",\n         \"metadata\": {\n            \"Odio ipsam est sit.\": \"Est possimus animi vel.\"\n         },\n         \"name\": \"Alias hic aliquid iste similique voluptatem magnam.\",\n         \"subject\": \"Nihil occaecati.\",\n         \"text_content\": \"Provident consequatur est aliquid reprehenderit.\"\n      }\n   }'")
		}
		if body.Template == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("template", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	var id string
	{
		id = emailUpdateTemplateID
	}
	var jwt *string
	{
		if emailUpdateTemplateJWT != "" {
			jwt = &emailUpdateTemplateJWT
		}
	}
	v := &email.UpdateTemplatePayload{}
	if body.Template != nil {
		v.Template = marshalUpdateEmailTemplateRequestRequestBodyToEmailUpdateEmailTemplateRequest(body.Template)
	}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildDeleteTemplatePayload builds the payload for the email delete_template
// endpoint from CLI flags.
func BuildDeleteTemplatePayload(emailDeleteTemplateID string, emailDeleteTemplateJWT string) (*email.DeleteTemplatePayload, error) {
	var id string
	{
		id = emailDeleteTemplateID
	}
	var jwt *string
	{
		if emailDeleteTemplateJWT != "" {
			jwt = &emailDeleteTemplateJWT
		}
	}
	v := &email.DeleteTemplatePayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildSendPayload builds the payload for the email send endpoint from CLI
// flags.
func BuildSendPayload(emailSendBody string, emailSendJWT string) (*email.SendPayload, error) {
	var err error
	var body SendRequestBody
	{
		err = json.Unmarshal([]byte(emailSendBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"bcc\": [\n         \"Consequuntur saepe earum.\",\n         \"Placeat minus aperiam.\",\n         \"Totam et ut et est.\",\n         \"Impedit eum id voluptas nisi.\"\n      ],\n      \"cc\": [\n         \"Neque aperiam quis praesentium voluptas iusto.\",\n         \"Eius ut perspiciatis.\",\n         \"Quia sunt aut necessitatibus at perferendis id.\",\n         \"Sunt architecto soluta et ab.\"\n      ],\n      \"from\": \"no-reply@example.com\",\n      \"headers\": {\n         \"Est culpa velit soluta qui.\": \"Sunt doloremque.\",\n         \"Molestiae inventore nostrum.\": \"Dicta eaque sint.\",\n         \"Quis sit ullam quia quo et.\": \"Qui corrupti.\"\n      },\n      \"html_content\": \"Nemo tenetur molestiae quas fugiat.\",\n      \"metadata\": {\n         \"Est nulla mollitia qui.\": \"Sapiente consequatur omnis omnis voluptatibus.\"\n      },\n      \"reply_to\": \"Sit quis sit distinctio non dolores.\",\n      \"subject\": \"Important information\",\n      \"text_content\": \"Adipisci sed animi dolor.\",\n      \"to\": [\n         \"user@example.com\"\n      ]\n   }'")
		}
		if body.To == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("to", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	var jwt *string
	{
		if emailSendJWT != "" {
			jwt = &emailSendJWT
		}
	}
	v := &email.SendPayload{
		From:        body.From,
		Subject:     body.Subject,
		HTMLContent: body.HTMLContent,
		TextContent: body.TextContent,
		ReplyTo:     body.ReplyTo,
	}
	if body.To != nil {
		v.To = make([]string, len(body.To))
		for i, val := range body.To {
			v.To[i] = val
		}
	} else {
		v.To = []string{}
	}
	if body.Cc != nil {
		v.Cc = make([]string, len(body.Cc))
		for i, val := range body.Cc {
			v.Cc[i] = val
		}
	}
	if body.Bcc != nil {
		v.Bcc = make([]string, len(body.Bcc))
		for i, val := range body.Bcc {
			v.Bcc[i] = val
		}
	}
	if body.Headers != nil {
		v.Headers = make(map[string]string, len(body.Headers))
		for key, val := range body.Headers {
			tk := key
			tv := val
			v.Headers[tk] = tv
		}
	}
	if body.Metadata != nil {
		v.Metadata = make(map[string]any, len(body.Metadata))
		for key, val := range body.Metadata {
			tk := key
			tv := val
			v.Metadata[tk] = tv
		}
	}
	v.JWT = jwt

	return v, nil
}

// BuildSendTemplatePayload builds the payload for the email send_template
// endpoint from CLI flags.
func BuildSendTemplatePayload(emailSendTemplateBody string, emailSendTemplateJWT string) (*email.SendTemplatePayload, error) {
	var err error
	var body SendTemplateRequestBody
	{
		err = json.Unmarshal([]byte(emailSendTemplateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"bcc\": [\n         \"Blanditiis atque et voluptatibus non beatae accusamus.\",\n         \"Doloribus neque non porro qui aut.\",\n         \"Eligendi laudantium impedit ut est aut.\",\n         \"Eos id itaque.\"\n      ],\n      \"cc\": [\n         \"Necessitatibus ut velit.\",\n         \"Assumenda aut quos.\",\n         \"Ex aut et repellat praesentium omnis.\"\n      ],\n      \"from\": \"no-reply@example.com\",\n      \"headers\": {\n         \"Ad sed sint.\": \"Non quibusdam ut excepturi quia enim modi.\",\n         \"Rerum dolor necessitatibus perferendis.\": \"In ratione.\"\n      },\n      \"locale\": \"Dignissimos debitis mollitia.\",\n      \"metadata\": {\n         \"Corrupti inventore.\": \"Architecto beatae quo.\",\n         \"Eos sunt.\": \"At qui.\",\n         \"Voluptas eveniet deleniti aliquam quam qui eaque.\": \"Earum excepturi voluptatem ab in consequuntur.\"\n      },\n      \"organization_id\": \"Quidem aut quisquam culpa sint.\",\n      \"reply_to\": \"Eos quia ullam quia dolores.\",\n      \"subject\": \"Quaerat exercitationem nihil et aut aliquam.\",\n      \"template_data\": {\n         \"name\": \"John Doe\"\n      },\n      \"template_type\": \"welcome\",\n      \"to\": [\n         \"user@example.com\"\n      ]\n   }'")
		}
		if body.To == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("to", "body"))
		}
		if body.TemplateData == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("template_data", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	var jwt *string
	{
		if emailSendTemplateJWT != "" {
			jwt = &emailSendTemplateJWT
		}
	}
	v := &email.SendTemplatePayload{
		From:           body.From,
		Subject:        body.Subject,
		TemplateType:   body.TemplateType,
		OrganizationID: body.OrganizationID,
		Locale:         body.Locale,
		ReplyTo:        body.ReplyTo,
	}
	if body.To != nil {
		v.To = make([]string, len(body.To))
		for i, val := range body.To {
			v.To[i] = val
		}
	} else {
		v.To = []string{}
	}
	if body.TemplateData != nil {
		v.TemplateData = make(map[string]any, len(body.TemplateData))
		for key, val := range body.TemplateData {
			tk := key
			tv := val
			v.TemplateData[tk] = tv
		}
	}
	{
		var zero string
		if v.Locale == zero {
			v.Locale = "en"
		}
	}
	if body.Cc != nil {
		v.Cc = make([]string, len(body.Cc))
		for i, val := range body.Cc {
			v.Cc[i] = val
		}
	}
	if body.Bcc != nil {
		v.Bcc = make([]string, len(body.Bcc))
		for i, val := range body.Bcc {
			v.Bcc[i] = val
		}
	}
	if body.Headers != nil {
		v.Headers = make(map[string]string, len(body.Headers))
		for key, val := range body.Headers {
			tk := key
			tv := val
			v.Headers[tk] = tv
		}
	}
	if body.Metadata != nil {
		v.Metadata = make(map[string]any, len(body.Metadata))
		for key, val := range body.Metadata {
			tk := key
			tv := val
			v.Metadata[tk] = tv
		}
	}
	v.JWT = jwt

	return v, nil
}
