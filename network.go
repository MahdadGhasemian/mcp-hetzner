package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// NetworkReadArgs represents the arguments required to read an Network by ID or Name.
// It contains the Network ID or Name that is needed to perform the lookup.
type NetworkReadArgs struct {
	IDOrName string `json:"id_or_name" jsonschema:"required,description=The network id or name to be searched"`
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
		Name:        "get_a_network_by_id_or_name",
		Description: "Retrieves a Network by its ID or Name. Get retrieves a network by its ID if the input can be parsed as an integer, otherwise it retrieves a network by its name. If the network does not exist, nil is returned.",
		Handler: func(args NetworkReadArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Network, error) {
				result, _, err := client.Network.Get(context.Background(), args.IDOrName)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
}
