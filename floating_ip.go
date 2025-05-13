package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// FloatingIPReadByIDArgs represents the arguments required to read an FloatingIP by ID.
// It contains the FloatingIP ID that is needed to perform the lookup.
type FloatingIPReadByIDArgs struct {
	ID int64 `json:"id" jsonschema:"required,description=The Floating IP id to be searched"`
}

// FloatingIPReadByNameArgs represents the arguments required to read an FloatingIP by Name.
// It contains the FloatingIP Name that is needed to perform the lookup.
type FloatingIPReadByNameArgs struct {
	Name string `json:"name" jsonschema:"required,description=The Floating IP name to be searched"`
}

// FloatingIPTools
var floatingIPTools = []Tool{
	{
		Name:        "get_all_floatingIPs",
		Description: "Returns all FloatingIPs objects.",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.FloatingIP, error) {
				result, err := client.FloatingIP.All(context.Background())
				return result, err
			})
		},
	},
	{
		Name:        "get_a_floatingIP_by_id",
		Description: "Retrieves a FloatingIP by its ID. If the FloatingIP does not exist, nil is returned.",
		Handler: func(args FloatingIPReadByIDArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.FloatingIP, error) {
				result, _, err := client.FloatingIP.GetByID(context.Background(), args.ID)
				return result, err
			})
		},
	},
	{
		Name:        "get_a_floatingIP_by_name",
		Description: "Retrieves a FloatingIP by its Name. If the FloatingIP does not exist, nil is returned.",
		Handler: func(args FloatingIPReadByNameArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.FloatingIP, error) {
				result, _, err := client.FloatingIP.GetByName(context.Background(), args.Name)
				return result, err
			})
		},
	},
}
