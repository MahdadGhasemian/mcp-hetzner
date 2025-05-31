package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// LoadBalancerTypeReadArgs represents the arguments required to read an LoadBalancerType by ID or Name.
// It contains the LoadBalancerType ID or Name that is needed to perform the lookup.
type LoadBalancerTypeReadArgs struct {
	IDOrName string `json:"id_or_name" jsonschema:"required,description=The Load Balancer Type id or name to be searched"`
}

// LoadBalancerTypeTools
var loadBalancerTypeTools = []Tool{
	{
		Name:        "get_all_load_balancer_types",
		Description: "Returns all LoadBalancerTypes objects.",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.LoadBalancerType, error) {
				result, err := client.LoadBalancerType.All(context.Background())
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_load_balancer_type_by_id_or_name",
		Description: "Retrieves a LoadBalancerType by its ID or Name. Get retrieves a load balancer type by its ID if the input can be parsed as an integer, otherwise it retrieves a load balancer type by its name. If the load balancer type does not exist, nil is returned.",
		Handler: func(args LoadBalancerTypeReadArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.LoadBalancerType, error) {
				result, _, err := client.LoadBalancerType.Get(context.Background(), args.IDOrName)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
}
