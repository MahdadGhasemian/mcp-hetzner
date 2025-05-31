package main

import (
	"context"
	"time"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// CertificateReadArgs represents the arguments to read an Certificate by ID or Name.
// It contains the Certificate ID or Name that is needed to perform the lookup.
type CertificateReadArgs struct {
	IDOrName string `json:"id_or_name" jsonschema:"required,description=The certificate id or name to be searched"`
}

type CertificateResponse struct {
	ID             int64  `json:"id" jsonschema:"required,description=Unique identifier of the certificate"`
	Name           string `json:"name" jsonschema:"required,description=The name of the certificate"`
	Labels         map[string]string
	Type           string    `json:"type" jsonschema:"description=The type of the certificate either of uploaded or managed"`
	Created        time.Time `json:"created" jsonschema:"description=Timestamp of when the certificate was created"`
	NotValidBefore time.Time `json:"not_valid_before" jsonschema:"description=Timestamp of when the certificate is not valid before"`
	NotValidAfter  time.Time `json:"not_valid_after" jsonschema:"description=Timestamp of when the certificate is not valid after"`
	DomainNames    []string  `json:"domain_names" jsonschema:"description=List of domain names for the certificate"`
	Fingerprint    string    `json:"fingerprint" jsonschema:"description=The fingerprint of the certificate"`
}

func toCertificateResponse(c *hcloud.Certificate) *CertificateResponse {
	if c == nil {
		return nil
	}
	return &CertificateResponse{
		ID:             c.ID,
		Name:           c.Name,
		Labels:         c.Labels,
		Type:           string(c.Type),
		Created:        c.Created,
		NotValidBefore: c.NotValidBefore,
		NotValidAfter:  c.NotValidAfter,
		DomainNames:    c.DomainNames,
		Fingerprint:    c.Fingerprint,
	}
}

// CertificateTools
var certificateTools = []Tool{
	{
		Name:        "get_all_certificates",
		Description: "Returns all Certificates objects.",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*CertificateResponse, error) {
				result, err := client.Certificate.All(context.Background())
				if err != nil {
					return nil, err
				}
				var filtered []*CertificateResponse
				for _, c := range result {
					filtered = append(filtered, toCertificateResponse(c))
				}
				return filtered, nil
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_certificate_by_id_or_name",
		Description: "Retrieves a Certificate by its ID or Name. Get retrieves a Certificate by its ID if the input can be parsed as an integer, otherwise it retrieves a Certificate by its name. If the Certificate does not exist, nil is returned.",
		Handler: func(args CertificateReadArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*CertificateResponse, error) {
				result, _, err := client.Certificate.Get(context.Background(), args.IDOrName)
				if err != nil {
					return nil, err
				}
				return toCertificateResponse(result), nil
			})
		},
		Restriction: RestrictionReadOnly,
	},
}
