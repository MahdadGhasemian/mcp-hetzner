package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// ImageReadByIDArgs represents the arguments required to read an Image by ID.
// It contains the Image ID that is needed to perform the lookup.
type ImageReadByIDArgs struct {
	ID int64 `json:"id" jsonschema:"required,description=The image id to be searched"`
}

// ImageTools
var imageTools = []Tool{
	{
		Name:        "get_all_images",
		Description: "Returns all Images objects.",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.Image, error) {
				result, err := client.Image.All(context.Background())
				return result, err
			})
		},
	},
	{
		Name:        "get_a_image_by_id",
		Description: "Retrieves a Image by its ID. If the Image does not exist, nil is returned.",
		Handler: func(args ImageReadByIDArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Image, error) {
				result, _, err := client.Image.GetByID(context.Background(), args.ID)
				return result, err
			})
		},
	},
}
