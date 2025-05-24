package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// ServerTypeReadByIDArgs represents the arguments required to read an ServerType by ID.
// It contains the ServerType ID that is needed to perform the lookup.
type ServerTypeReadByIDArgs struct {
	ID int64 `json:"id" jsonschema:"required,description=The Server Type id to be searched"`
}

// ServerTypeReadByNameArgs represents the arguments required to read an ServerType by Name.
// It contains the ServerType Name that is needed to perform the lookup.
type ServerTypeReadByNameArgs struct {
	Name string `json:"name" jsonschema:"required,description=The Server Type name to be searched"`
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
		Name:        "get_a_server_type_by_id",
		Description: "Retrieves a ServerType by its ID. If the ServerType does not exist, nil is returned.",
		Handler: func(args ServerTypeReadByIDArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.ServerType, error) {
				result, _, err := client.ServerType.GetByID(context.Background(), args.ID)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_server_type_by_name",
		Description: "Retrieves a ServerType by its Name. If the ServerType does not exist, nil is returned.",
		Handler: func(args ServerTypeReadByNameArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.ServerType, error) {
				result, _, err := client.ServerType.GetByName(context.Background(), args.Name)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
}
