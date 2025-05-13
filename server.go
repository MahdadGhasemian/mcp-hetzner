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

// ServerCreateArgs contains the necessary fields to create a new server,
type ServerCreateArgs struct {
	Name           string            `json:"name" jsonschema:"required"`
	ServerTypeName string            `json:"server_type_name" jsonschema:"required"`
	ImageName      string            `json:"image_name" jsonschema:"required"`
	LocationName   string            `json:"location_name" jsonschema:"required"`
	DatacenterName string            `json:"datacenter_name" jsonschema:"required"`
	SSHKeyNames    []string          `json:"ssh_key_names" jsonschema:"required"`
	Labels         map[string]string `json:"labels,omitempty"`
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
		Name:        "create_a_server",
		Description: "Creates a new Server. Returns preliminary information about the Server as well as an Action that covers progress of creation.",
		Handler: func(args ServerCreateArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (hcloud.ServerCreateResult, error) {
				result, _, err := client.Server.Create(context.Background(), hcloud.ServerCreateOpts{
					Name:   args.Name,
					Labels: args.Labels,
				})
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
