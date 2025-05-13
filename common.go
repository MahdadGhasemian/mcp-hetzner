package main

// EmptyString is a constant that represents an empty string.
const EmptyString = ""

// NoArgs represents an empty structure used when no arguments are required.
type NoArgs struct{}

// ListArgs represents the arguments for listing resources.
// It includes pagination information like page number, items per page, and a label selector for filtering.
type ListArgs struct {
	Page          int    `json:"page" jsonschema:"description=Page (starting at 1)"`
	PerPage       int    `json:"per_page" jsonschema:"description=Items per page (0 means default)"`
	LabelSelector string `json:"label_selector" jsonschema:"description=Label selector for filtering by labels"`
}

// Tool represents a tool with a name, description, and handler function.
type Tool struct {
	Name        string
	Description string
	Handler     any
}
