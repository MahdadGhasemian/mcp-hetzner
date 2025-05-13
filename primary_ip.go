package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// PrimaryIPReadByIDArgs represents the arguments required to read an PrimaryIP by ID.
// It contains the PrimaryIP ID that is needed to perform the lookup.
type PrimaryIPReadByIDArgs struct {
	ID int64 `json:"id" jsonschema:"required,description=The Primary IP id to be searched"`
}

// PrimaryIPReadByNameArgs represents the arguments required to read an PrimaryIP by Name.
// It contains the PrimaryIP Name that is needed to perform the lookup.
type PrimaryIPReadByNameArgs struct {
	Name string `json:"name" jsonschema:"required,description=The Primary IP name to be searched"`
}

// PrimaryIPReadByIPArgs represents the arguments required to read an PrimaryIP by IP.
// It contains the PrimaryIP IP that is needed to perform the lookup.
type PrimaryIPReadByIPArgs struct {
	IP string `json:"ip" jsonschema:"required,description=The Primary IP ip to be searched"`
}

// PrimaryIPTools
var primaryIPTools = []Tool{
	{
		Name:        "get_all_primary_ips",
		Description: "Returns all PrimaryIPs objects.",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.PrimaryIP, error) {
				result, err := client.PrimaryIP.All(context.Background())
				return result, err
			})
		},
	},
	{
		Name:        "get_a_primary_ip_by_id",
		Description: "Retrieves a PrimaryIP by its ID. If the PrimaryIP does not exist, nil is returned.",
		Handler: func(args PrimaryIPReadByIDArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.PrimaryIP, error) {
				result, _, err := client.PrimaryIP.GetByID(context.Background(), args.ID)
				return result, err
			})
		},
	},
	{
		Name:        "get_a_primary_ip_by_name",
		Description: "Retrieves a PrimaryIP by its Name. If the PrimaryIP does not exist, nil is returned.",
		Handler: func(args PrimaryIPReadByNameArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.PrimaryIP, error) {
				result, _, err := client.PrimaryIP.GetByName(context.Background(), args.Name)
				return result, err
			})
		},
	},
	{
		Name:        "get_a_primary_ip_by_ip",
		Description: "Retrieves a PrimaryIP by its IP. If the PrimaryIP does not exist, nil is returned.",
		Handler: func(args PrimaryIPReadByIPArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.PrimaryIP, error) {
				result, _, err := client.PrimaryIP.GetByIP(context.Background(), args.IP)
				return result, err
			})
		},
	},
}
