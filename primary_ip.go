package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// PrimaryIPReadArgs represents the arguments required to read an PrimaryIP by ID or Name.
// It contains the PrimaryIP ID or Name that is needed to perform the lookup.
type PrimaryIPReadArgs struct {
	IDOrName string `json:"id_or_name" jsonschema:"required,description=The Primary IP id or name to be searched"`
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
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_primary_ip_by_id_or_name",
		Description: "Retrieves a PrimaryIP by its ID or Name. Get retrieves a Primary IP by its ID if the input can be parsed as an integer, otherwise it retrieves a Primary IP by its name. If the Primary IP does not exist, nil is returned.",
		Handler: func(args PrimaryIPReadArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.PrimaryIP, error) {
				result, _, err := client.PrimaryIP.Get(context.Background(), args.IDOrName)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
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
		Restriction: RestrictionReadOnly,
	},
}
