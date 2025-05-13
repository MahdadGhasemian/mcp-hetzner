package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// ServerReadArgs represents the arguments required to read a server.
// It contains the server ID that is needed to perform the lookup.
type ServerReadArgs struct {
	ServerID int64 `json:"server_id" jsonschema:"required,description=The server id to be searched"`
}

// ServerTools
var serverTools = []Tool{
	{
		Name:        "get_server_list",
		Description: "Returns all existing Server objects",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.Server, error) {
				result, _, err := client.Server.List(context.Background(), hcloud.ServerListOpts{})
				return result, err
			})
		},
	},
	{
		Name:        "get_server_info_by_id",
		Description: "Get a server by its ID, it returns the server object info",
		Handler: func(args ServerReadArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Server, error) {
				result, _, err := client.Server.GetByID(context.Background(), args.ServerID)
				return result, err
			})
		},
	},
}
