package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// ServerReadByIDArgs represents the arguments required to read an Server by ID.
// It contains the Server ID that is needed to perform the lookup.
type ServerReadByIDArgs struct {
	ID int64 `json:"id" jsonschema:"required,description=The server id to be searched"`
}

// ServerReadByNameArgs represents the arguments required to read an Server by Name.
// It contains the Server Name that is needed to perform the lookup.
type ServerReadByNameArgs struct {
	Name string `json:"name" jsonschema:"required,description=The server name to be searched"`
}

// ServerTools
var serverTools = []Tool{
	{
		Name:        "get_all_servers",
		Description: "Returns all Servers objects.",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.Server, error) {
				result, err := client.Server.All(context.Background())
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_server_by_id",
		Description: "Retrieves a Server by its ID. If the Server does not exist, nil is returned.",
		Handler: func(args ServerReadByIDArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Server, error) {
				result, _, err := client.Server.GetByID(context.Background(), args.ID)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_server_by_name",
		Description: "Retrieves a Server by its Name. If the Server does not exist, nil is returned.",
		Handler: func(args ServerReadByNameArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Server, error) {
				result, _, err := client.Server.GetByName(context.Background(), args.Name)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
}
