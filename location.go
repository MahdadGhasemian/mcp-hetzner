package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// LocationReadArgs represents the arguments required to read an Location by ID or Name.
// It contains the Location ID or Name that is needed to perform the lookup.
type LocationReadArgs struct {
	IDOrName string `json:"id_or_name" jsonschema:"required,description=The location id or name to be searched"`
}

// LocationTools
var locationTools = []Tool{
	{
		Name:        "get_all_locations",
		Description: "Returns all Locations objects.",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.Location, error) {
				result, err := client.Location.All(context.Background())
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_location_by_id_or_name",
		Description: "Retrieves a Location by its ID or Name, Get retrieves a Location by its ID if the input can be parsed as an integer, otherwise it retrieves a Location by its name. If the Location does not exist, nil is returned.",
		Handler: func(args LocationReadArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Location, error) {
				result, _, err := client.Location.Get(context.Background(), args.IDOrName)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
}
