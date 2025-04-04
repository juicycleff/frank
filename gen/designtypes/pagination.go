// Code generated by goa v3.20.0, DO NOT EDIT.
//
// User types
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package designtypes

// Pagination parameters
type Pagination struct {
	// Offset
	Offset int `json:"offset"`
	// Limit
	Limit int `json:"limit"`
	// Total number of items
	Total int `json:"total"`
	// Total number of pages
	TotalPages int `json:"total_pages,totalPages"`
	// Current page number
	CurrentPage int `json:"current_page,currentPage"`
	// Has next page
	HasNext bool `json:"has_next,hasNext"`
	// Has previous page
	HasPrevious bool `json:"has_previous,hasPrevious"`
}
