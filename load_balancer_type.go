package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// LoadBalancerTypeReadByIDArgs represents the arguments required to read an LoadBalancerType by ID.
// It contains the LoadBalancerType ID that is needed to perform the lookup.
type LoadBalancerTypeReadByIDArgs struct {
	ID int64 `json:"id" jsonschema:"required,description=The Load Balancer Type id to be searched"`
}

// LoadBalancerTypeReadByNameArgs represents the arguments required to read an LoadBalancerType by Name.
// It contains the LoadBalancerType Name that is needed to perform the lookup.
type LoadBalancerTypeReadByNameArgs struct {
	Name string `json:"name" jsonschema:"required,description=The Load Balancer Type name to be searched"`
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
	},
	{
		Name:        "get_a_load_balancer_type_by_id",
		Description: "Retrieves a LoadBalancerType by its ID. If the LoadBalancerType does not exist, nil is returned.",
		Handler: func(args LoadBalancerTypeReadByIDArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.LoadBalancerType, error) {
				result, _, err := client.LoadBalancerType.GetByID(context.Background(), args.ID)
				return result, err
			})
		},
	},
	{
		Name:        "get_a_load_balancer_type_by_name",
		Description: "Retrieves a LoadBalancerType by its Name. If the LoadBalancerType does not exist, nil is returned.",
		Handler: func(args LoadBalancerTypeReadByNameArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.LoadBalancerType, error) {
				result, _, err := client.LoadBalancerType.GetByName(context.Background(), args.Name)
				return result, err
			})
		},
	},
}
