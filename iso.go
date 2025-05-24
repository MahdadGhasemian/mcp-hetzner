package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// ISOReadByIDArgs represents the arguments required to read an ISO by ID.
// It contains the ISO ID that is needed to perform the lookup.
type ISOReadByIDArgs struct {
	ID int64 `json:"id" jsonschema:"required,description=The ISO id to be searched"`
}

// ISOReadByNameArgs represents the arguments required to read an ISO by Name.
// It contains the ISO Name that is needed to perform the lookup.
type ISOReadByNameArgs struct {
	Name string `json:"name" jsonschema:"required,description=The ISO name to be searched"`
}

// ISOTools
var isoTools = []Tool{
	{
		Name:        "get_all_isos",
		Description: "Returns all ISOs objects.",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.ISO, error) {
				result, err := client.ISO.All(context.Background())
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_iso_by_id",
		Description: "Retrieves a ISO by its ID. If the ISO does not exist, nil is returned.",
		Handler: func(args ISOReadByIDArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.ISO, error) {
				result, _, err := client.ISO.GetByID(context.Background(), args.ID)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_iso_by_name",
		Description: "Retrieves a ISO by its Name. If the ISO does not exist, nil is returned.",
		Handler: func(args ISOReadByNameArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.ISO, error) {
				result, _, err := client.ISO.GetByName(context.Background(), args.Name)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
}
