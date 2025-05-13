package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// FirewallReadArgs represents the arguments required to fetch a specific firewall by its ID.
type FirewallReadArgs struct {
	FirewallID int64 `json:"firewall_id" jsonschema:"required,description=The firewall id to be searched"`
}

// FirewallReadListArgs defines the filtering and sorting options for listing firewalls.
type FirewallReadListArgs struct {
	ListArgs
	Name string   `json:"name" jsonschema:"description=Filter resources by their name."`
	Sort []string `json:"sort" jsonschema:"description=Sort fields"`
}

// FirewallTools
var firewallTools = []Tool{
	{
		Name:        "get_firewall_list",
		Description: "Returns all existing Firewall objects",
		Handler: func(args FirewallReadListArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.Firewall, error) {
				result, _, err := client.Firewall.List(context.Background(), hcloud.FirewallListOpts{
					ListOpts: hcloud.ListOpts(args.ListArgs),
					Name:     args.Name,
					Sort:     args.Sort,
				})
				return result, err
			})
		},
	},
	{
		Name:        "get_firewall_by_id",
		Description: "Get a Firewall by its ID, it returns a specific Firewall object info",
		Handler: func(args FirewallReadArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Firewall, error) {
				result, _, err := client.Firewall.GetByID(context.Background(), args.FirewallID)
				return result, err
			})
		},
	},
}
