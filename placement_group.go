package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// PlacementGroupReadArgs represents the arguments required to read an PlacementGroup by ID or Name.
// It contains the PlacementGroup ID or Name that is needed to perform the lookup.
type PlacementGroupReadArgs struct {
	IDOrName string `json:"id_or_name" jsonschema:"required,description=The Placement Group id or name to be searched"`
}

// PlacementGroupTools
var placementGroupTools = []Tool{
	{
		Name:        "get_all_placement_groups",
		Description: "Returns all PlacementGroups objects.",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.PlacementGroup, error) {
				result, err := client.PlacementGroup.All(context.Background())
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_placement_group_by_id_or_name",
		Description: "Retrieves a PlacementGroup by its ID or Name.",
		Handler: func(args PlacementGroupReadArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.PlacementGroup, error) {
				result, _, err := client.PlacementGroup.Get(context.Background(), args.IDOrName)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
}
