package main

import (
	"context"
	"time"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// SSHKeyReadByIDArgs represents the arguments required to read an SSH key by ID.
// It contains the SSH key ID that is needed to perform the lookup.
type SSHKeyReadByIDArgs struct {
	ID int64 `json:"id" jsonschema:"required,description=The ssh key id to be searched"`
}

// SSHKeyReadByNameArgs represents the arguments required to read an SSH key by Name.
// It contains the SSH key Name that is needed to perform the lookup.
type SSHKeyReadByNameArgs struct {
	Name string `json:"name" jsonschema:"required,description=The ssh key name to be searched"`
}

// SSHKeyReadByFingerprintArgs represents the arguments required to read an SSH key by Fingerprint.
// It contains the SSH key Fingerprint that is needed to perform the lookup.
type SSHKeyReadByFingerprintArgs struct {
	Fingerprint string `json:"fingerprint" jsonschema:"required,description=The ssh key fingerprint to be searched"`
}

type SSHKeyResponse struct {
	ID          int64             `json:"id" jsonschema:"required,description=Unique identifier of the ssh key"`
	Name        string            `json:"name" jsonschema:"required,description=The name of the ssh key"`
	Fingerprint string            `json:"fingerprint" jsonschema:"required,description=The fingerprint of the ssh key"`
	PublicKey   string            `json:"public_key" jsonschema:"required,description=The public key of the ssh key"`
	Labels      map[string]string `json:"labels" jsonschema:"description=User-defined labels for the ssh key"`
	Created     time.Time         `json:"created" jsonschema:"required,description=Timestamp of when the ssh key was created"`
}

func tossSSHKeyResponse(s *hcloud.SSHKey) *SSHKeyResponse {
	if s == nil {
		return nil
	}

	return &SSHKeyResponse{
		ID:          s.ID,
		Name:        s.Name,
		Fingerprint: s.Fingerprint,
		PublicKey:   s.PublicKey,
		Labels:      s.Labels,
		Created:     s.Created,
	}
}

// SSHKeyTools
var sshkeyTools = []Tool{
	{
		Name:        "get_all_ssh_keys",
		Description: "Returns all ssh-key objects. SSH keys are public keys you provide to the cloud system. They can be injected into Servers at creation time.",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*SSHKeyResponse, error) {
				result, err := client.SSHKey.All(context.Background())
				if err != nil {
					return nil, err
				}
				var filtered []*SSHKeyResponse
				for _, s := range result {
					filtered = append(filtered, tossSSHKeyResponse(s))
				}
				return filtered, nil
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_ssh_key_by_id",
		Description: "Retrieves a SSH key by its ID, it returns a specific ssh key object info. If the SSH key does not exist, nil is returned.",
		Handler: func(args SSHKeyReadByIDArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*SSHKeyResponse, error) {
				result, _, err := client.SSHKey.GetByID(context.Background(), args.ID)
				return tossSSHKeyResponse(result), err
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_ssh_key_by_name",
		Description: "Retrieves a SSH key by its Name. If the SSH key does not exist, nil is returned.",
		Handler: func(args SSHKeyReadByNameArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*SSHKeyResponse, error) {
				result, _, err := client.SSHKey.GetByName(context.Background(), args.Name)
				return tossSSHKeyResponse(result), err
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_ssh_key_by_fingerprint",
		Description: "Retrieves a SSH key by its Fingerprint. If the SSH key does not exist, nil is returned.",
		Handler: func(args SSHKeyReadByFingerprintArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*SSHKeyResponse, error) {
				result, _, err := client.SSHKey.GetByName(context.Background(), args.Fingerprint)
				return tossSSHKeyResponse(result), err
			})
		},
		Restriction: RestrictionReadOnly,
	},
}
