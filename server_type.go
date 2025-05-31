package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// ServerTypeReadArgs represents the arguments required to read an ServerType by ID or Name.
// It contains the ServerType ID or Name that is needed to perform the lookup.
type ServerTypeReadArgs struct {
	IDOrName string `json:"id_or_name" jsonschema:"required,description=The Server Type id or name to be searched"`
}

// ServerTypeTools
var serverTypeTools = []Tool{
	{
		Name:        "get_all_server_types",
		Description: "Returns all ServerTypes objects.",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.ServerType, error) {
				result, err := client.ServerType.All(context.Background())
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_server_type_by_id_or_name",
		Description: "Retrieves a ServerType by its ID or Name. Get retrieves a server type by its ID if the input can be parsed as an integer, otherwise it retrieves a server type by its name. If the server type does not exist, nil is returned.",
		Handler: func(args ServerTypeReadArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.ServerType, error) {
				result, _, err := client.ServerType.Get(context.Background(), args.IDOrName)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
}
