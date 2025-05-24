package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// VolumeReadByIDArgs represents the arguments required to read an Volume by ID.
// It contains the Volume ID that is needed to perform the lookup.
type VolumeReadByIDArgs struct {
	ID int64 `json:"id" jsonschema:"required,description=The volume id to be searched"`
}

// VolumeReadByNameArgs represents the arguments required to read an Volume by Name.
// It contains the Volume Name that is needed to perform the lookup.
type VolumeReadByNameArgs struct {
	Name string `json:"name" jsonschema:"required,description=The volume name to be searched"`
}

// VolumeTools
var volumeTools = []Tool{
	{
		Name:        "get_all_volumes",
		Description: "Returns all Volumes objects.",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.Volume, error) {
				result, err := client.Volume.All(context.Background())
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_volume_by_id",
		Description: "Retrieves a Volume by its ID. If the Volume does not exist, nil is returned.",
		Handler: func(args VolumeReadByIDArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Volume, error) {
				result, _, err := client.Volume.GetByID(context.Background(), args.ID)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_volume_by_name",
		Description: "Retrieves a Volume by its Name. If the Volume does not exist, nil is returned.",
		Handler: func(args VolumeReadByNameArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Volume, error) {
				result, _, err := client.Volume.GetByName(context.Background(), args.Name)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
}
