package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// ISOReadArgs represents the arguments required to read an ISO by ID or Name.
// It contains the ISO ID or Name that is needed to perform the lookup.
type ISOReadArgs struct {
	IDOrName string `json:"id_or_name" jsonschema:"required,description=The ISO id or name to be searched"`
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
		Name:        "get_a_iso_by_id_or_name",
		Description: "Retrieves a ISO by its ID or Name.",
		Handler: func(args ISOReadArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.ISO, error) {
				result, _, err := client.ISO.Get(context.Background(), args.IDOrName)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
}
