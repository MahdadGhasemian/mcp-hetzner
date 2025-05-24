package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// NetworkReadByIDArgs represents the arguments required to read an Network by ID.
// It contains the Network ID that is needed to perform the lookup.
type NetworkReadByIDArgs struct {
	ID int64 `json:"id" jsonschema:"required,description=The network id to be searched"`
}

// NetworkReadByNameArgs represents the arguments required to read an Network by Name.
// It contains the Network Name that is needed to perform the lookup.
type NetworkReadByNameArgs struct {
	Name string `json:"name" jsonschema:"required,description=The network name to be searched"`
}

// NetworkTools
var networkTools = []Tool{
	{
		Name:        "get_all_networks",
		Description: "Returns all Networks objects.",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.Network, error) {
				result, err := client.Network.All(context.Background())
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_network_by_id",
		Description: "Retrieves a Network by its ID. If the Network does not exist, nil is returned.",
		Handler: func(args NetworkReadByIDArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Network, error) {
				result, _, err := client.Network.GetByID(context.Background(), args.ID)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_network_by_name",
		Description: "Retrieves a Network by its Name. If the Network does not exist, nil is returned.",
		Handler: func(args NetworkReadByNameArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Network, error) {
				result, _, err := client.Network.GetByName(context.Background(), args.Name)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
}
