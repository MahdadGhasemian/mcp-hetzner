package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// FloatingIPReadArgs represents the arguments required to read an FloatingIP by ID or Name.
// It contains the FloatingIP ID or Name that is needed to perform the lookup.
type FloatingIPReadArgs struct {
	IDOrName string `json:"id_or_name" jsonschema:"required,description=The Floating IP id or name to be searched"`
}

// FloatingIPTools
var floatingIPTools = []Tool{
	{
		Name:        "get_all_floating_ips",
		Description: "Returns all FloatingIPs objects.",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.FloatingIP, error) {
				result, err := client.FloatingIP.All(context.Background())
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_floating_ip_by_id_or_name",
		Description: "Retrieves a FloatingIP by its ID or Name, Get retrieves a FloatingIP by its ID if the input can be parsed as an integer, otherwise it retrieves a FloatingIP by its name. If the FloatingIP does not exist, nil is returned.",
		Handler: func(args FloatingIPReadArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.FloatingIP, error) {
				result, _, err := client.FloatingIP.Get(context.Background(), args.IDOrName)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
}
