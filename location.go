package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// LocationReadArgs represents the arguments required to read a location.
// It contains the location ID that is needed to perform the lookup.
type LocationReadArgs struct {
	LocationID int64 `json:"location_id" jsonschema:"required,description=The location id to be searched"`
}

// LocationTools
var locationTools = []Tool{
	{
		Name:        "get_location_list",
		Description: "Returns all locations objects",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.Location, error) {
				result, _, err := client.Location.List(context.Background(), hcloud.LocationListOpts{})
				return result, err
			})
		},
	},
	{
		Name:        "get_location_info_by_id",
		Description: "Get a location by its ID, it returns the location object info",
		Handler: func(args LocationReadArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Location, error) {
				result, _, err := client.Location.GetByID(context.Background(), args.LocationID)
				return result, err
			})
		},
	},
}
