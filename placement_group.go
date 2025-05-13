package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// PlacementGroupReadByIDArgs represents the arguments required to read an PlacementGroup by ID.
// It contains the PlacementGroup ID that is needed to perform the lookup.
type PlacementGroupReadByIDArgs struct {
	ID int64 `json:"id" jsonschema:"required,description=The Placement Group id to be searched"`
}

// PlacementGroupReadByNameArgs represents the arguments required to read an PlacementGroup by Name.
// It contains the PlacementGroup Name that is needed to perform the lookup.
type PlacementGroupReadByNameArgs struct {
	Name string `json:"name" jsonschema:"required,description=The Placement Group name to be searched"`
}

// PlacementGroupTools
var placementGroupTools = []Tool{
	{
		Name:        "get_all_placementGroups",
		Description: "Returns all PlacementGroups objects.",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.PlacementGroup, error) {
				result, err := client.PlacementGroup.All(context.Background())
				return result, err
			})
		},
	},
	{
		Name:        "get_a_placementGroup_by_id",
		Description: "Retrieves a PlacementGroup by its ID. If the PlacementGroup does not exist, nil is returned.",
		Handler: func(args PlacementGroupReadByIDArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.PlacementGroup, error) {
				result, _, err := client.PlacementGroup.GetByID(context.Background(), args.ID)
				return result, err
			})
		},
	},
	{
		Name:        "get_a_placementGroup_by_name",
		Description: "Retrieves a PlacementGroup by its Name. If the PlacementGroup does not exist, nil is returned.",
		Handler: func(args PlacementGroupReadByNameArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.PlacementGroup, error) {
				result, _, err := client.PlacementGroup.GetByName(context.Background(), args.Name)
				return result, err
			})
		},
	},
}
