package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// LocationReadByIDArgs represents the arguments required to read an Location by ID.
// It contains the Location ID that is needed to perform the lookup.
type LocationReadByIDArgs struct {
	ID int64 `json:"id" jsonschema:"required,description=The location id to be searched"`
}

// LocationReadByNameArgs represents the arguments required to read an Location by Name.
// It contains the Location Name that is needed to perform the lookup.
type LocationReadByNameArgs struct {
	Name string `json:"name" jsonschema:"required,description=The location name to be searched"`
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
		Name:        "get_a_location_by_id",
		Description: "Retrieves a Location by its ID. If the Location does not exist, nil is returned.",
		Handler: func(args LocationReadByIDArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Location, error) {
				result, _, err := client.Location.GetByID(context.Background(), args.ID)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_location_by_name",
		Description: "Retrieves a Location by its Name. If the Location does not exist, nil is returned.",
		Handler: func(args LocationReadByNameArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Location, error) {
				result, _, err := client.Location.GetByName(context.Background(), args.Name)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
}
