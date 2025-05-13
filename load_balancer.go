package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// LoadBalancerReadByIDArgs represents the arguments required to read an LoadBalancer by ID.
// It contains the LoadBalancer ID that is needed to perform the lookup.
type LoadBalancerReadByIDArgs struct {
	ID int64 `json:"id" jsonschema:"required,description=The Load Balancer id to be searched"`
}

// LoadBalancerReadByNameArgs represents the arguments required to read an LoadBalancer by Name.
// It contains the LoadBalancer Name that is needed to perform the lookup.
type LoadBalancerReadByNameArgs struct {
	Name string `json:"name" jsonschema:"required,description=The Load Balancer name to be searched"`
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
	},
	{
		Name:        "get_a_load_balancer_by_id",
		Description: "Retrieves a LoadBalancer by its ID. If the LoadBalancer does not exist, nil is returned.",
		Handler: func(args LoadBalancerReadByIDArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.LoadBalancer, error) {
				result, _, err := client.LoadBalancer.GetByID(context.Background(), args.ID)
				return result, err
			})
		},
	},
	{
		Name:        "get_a_load_balancer_by_name",
		Description: "Retrieves a LoadBalancer by its Name. If the LoadBalancer does not exist, nil is returned.",
		Handler: func(args LoadBalancerReadByNameArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.LoadBalancer, error) {
				result, _, err := client.LoadBalancer.GetByName(context.Background(), args.Name)
				return result, err
			})
		},
	},
}
