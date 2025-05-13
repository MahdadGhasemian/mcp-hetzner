package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// CertificateReadByIDArgs represents the arguments required to read an Certificate by ID.
// It contains the Certificate ID that is needed to perform the lookup.
type CertificateReadByIDArgs struct {
	ID int64 `json:"id" jsonschema:"required,description=The certificate id to be searched"`
}

// CertificateReadByNameArgs represents the arguments required to read an Certificate by Name.
// It contains the Certificate Name that is needed to perform the lookup.
type CertificateReadByNameArgs struct {
	Name string `json:"name" jsonschema:"required,description=The certificate name to be searched"`
}

// CertificateTools
var certificateTools = []Tool{
	{
		Name:        "get_certificate_list",
		Description: "Returns all Certificates objects.",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.Certificate, error) {
				result, err := client.Certificate.All(context.Background())
				return result, err
			})
		},
	},
	{
		Name:        "get_certificate_by_id",
		Description: "Retrieves a Certificate by its ID. If the Certificate does not exist, nil is returned.",
		Handler: func(args CertificateReadByIDArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Certificate, error) {
				result, _, err := client.Certificate.GetByID(context.Background(), args.ID)
				return result, err
			})
		},
	},
	{
		Name:        "get_certificate_by_name",
		Description: "Retrieves a Certificate by its Name. If the Certificate does not exist, nil is returned.",
		Handler: func(args CertificateReadByNameArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Certificate, error) {
				result, _, err := client.Certificate.GetByName(context.Background(), args.Name)
				return result, err
			})
		},
	},
}
