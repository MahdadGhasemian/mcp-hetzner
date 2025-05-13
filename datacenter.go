package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// DatacenterReadByIDArgs represents the arguments required to read an Datacenter by ID.
// It contains the Datacenter ID that is needed to perform the lookup.
type DatacenterReadByIDArgs struct {
	ID int64 `json:"id" jsonschema:"required,description=The datacenter id to be searched"`
}

// DatacenterReadByNameArgs represents the arguments required to read an Datacenter by Name.
// It contains the Datacenter Name that is needed to perform the lookup.
type DatacenterReadByNameArgs struct {
	Name string `json:"name" jsonschema:"required,description=The datacenter name to be searched"`
}

// DatacenterTools
var datacenterTools = []Tool{
	{
		Name:        "get_all_datacenters",
		Description: "Returns all Datacenters objects.",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.Datacenter, error) {
				result, err := client.Datacenter.All(context.Background())
				return result, err
			})
		},
	},
	{
		Name:        "get_a_datacenter_by_id",
		Description: "Retrieves a Datacenter by its ID. If the Datacenter does not exist, nil is returned.",
		Handler: func(args DatacenterReadByIDArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Datacenter, error) {
				result, _, err := client.Datacenter.GetByID(context.Background(), args.ID)
				return result, err
			})
		},
	},
	{
		Name:        "get_a_datacenter_by_name",
		Description: "Retrieves a Datacenter by its Name. If the Datacenter does not exist, nil is returned.",
		Handler: func(args DatacenterReadByNameArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Datacenter, error) {
				result, _, err := client.Datacenter.GetByName(context.Background(), args.Name)
				return result, err
			})
		},
	},
}
