package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// FirewallReadByIDArgs represents the arguments required to read an Firewall by ID.
// It contains the Firewall ID that is needed to perform the lookup.
type FirewallReadByIDArgs struct {
	ID int64 `json:"id" jsonschema:"required,description=The firewall id to be searched"`
}

// FirewallReadByNameArgs represents the arguments required to read an Firewall by Name.
// It contains the Firewall Name that is needed to perform the lookup.
type FirewallReadByNameArgs struct {
	Name string `json:"name" jsonschema:"required,description=The firewall name to be searched"`
}

// FirewallTools
var firewallTools = []Tool{
	{
		Name:        "get_all_firewalls",
		Description: "Returns all Firewalls objects.",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.Firewall, error) {
				result, err := client.Firewall.All(context.Background())
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_firewall_by_id",
		Description: "Retrieves a Firewall by its ID. If the Firewall does not exist, nil is returned.",
		Handler: func(args FirewallReadByIDArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Firewall, error) {
				result, _, err := client.Firewall.GetByID(context.Background(), args.ID)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_firewall_by_name",
		Description: "Retrieves a Firewall by its Name. If the Firewall does not exist, nil is returned.",
		Handler: func(args FirewallReadByNameArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Firewall, error) {
				result, _, err := client.Firewall.GetByName(context.Background(), args.Name)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
}
