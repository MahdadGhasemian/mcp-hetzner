package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// DatacenterReadArgs represents the arguments required to read a datacenter.
// It contains the datacenter ID that is needed to perform the lookup.
type DatacenterReadArgs struct {
	DatacenterID int64 `json:"datacenter_id" jsonschema:"required,description=The datacenter id to be searched"`
}

// DatacenterTools
var datacenterTools = []Tool{
	{
		Name:        "get_datacenter_list",
		Description: "Returns all datacenters objects",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.Datacenter, error) {
				result, _, err := client.Datacenter.List(context.Background(), hcloud.DatacenterListOpts{})
				return result, err
			})
		},
	},
	{
		Name:        "get_datacenter_info_by_id",
		Description: "Get a datacenter by its ID, it returns the datacenter object info",
		Handler: func(args DatacenterReadArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Datacenter, error) {
				result, _, err := client.Datacenter.GetByID(context.Background(), args.DatacenterID)
				return result, err
			})
		},
	},
}
