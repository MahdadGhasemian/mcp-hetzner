package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// LoadBalancerReadArgs represents the arguments required to read an LoadBalancer by ID or Name.
// It contains the LoadBalancer ID or Name that is needed to perform the lookup.
type LoadBalancerReadArgs struct {
	IDOrName string `json:"id_or_name" jsonschema:"required,description=The Load Balancer id or name to be searched"`
}

// LoadBalancerTools
var loadBalancerTools = []Tool{
	{
		Name:        "get_all_load_balancers",
		Description: "Returns all LoadBalancers objects.",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.LoadBalancer, error) {
				result, err := client.LoadBalancer.All(context.Background())
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_load_balancer_by_id_or_name",
		Description: "Retrieves a LoadBalancer by its ID or Name. Get retrieves a load balancer by its ID if the input can be parsed as an integer, otherwise it retrieves a load balancer by its name. If the load balancer does not exist, nil is returned.",
		Handler: func(args LoadBalancerReadArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.LoadBalancer, error) {
				result, _, err := client.LoadBalancer.Get(context.Background(), args.IDOrName)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
}
