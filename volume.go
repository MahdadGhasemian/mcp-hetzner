package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// VolumeReadArgs represents the arguments required to read an Volume by ID or Name.
// It contains the Volume ID or Name that is needed to perform the lookup.
type VolumeReadArgs struct {
	IDOrName string `json:"id_or_name" jsonschema:"required,description=The volume id or name to be searched"`
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
		Name:        "get_a_volume_by_id_or_name",
		Description: "Retrieves a Volume by its ID or Name. Get retrieves a volume by its ID if the input can be parsed as an integer, otherwise it retrieves a volume by its name. If the volume does not exist, nil is returned.",
		Handler: func(args VolumeReadArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Volume, error) {
				result, _, err := client.Volume.Get(context.Background(), args.IDOrName)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
}
