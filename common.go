package main

// EmptyString is a constant that represents an empty string.
const EmptyString = ""

// NoArgs represents an empty structure used when no arguments are required.
type NoArgs struct{}

// Restriction represents the restriction of a tool.
type Restriction string

// RestrictionReadOnly represents a read-only restriction.
// It means that the tool can only read the resources.
const RestrictionReadOnly Restriction = "read_only"

// RestrictionReadWrite represents a read-write restriction.
// It means that the tool can read and write the resources.
const RestrictionReadWrite Restriction = "read_write"

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
	Restriction Restriction
}
