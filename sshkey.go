package main

import (
	"context"
	"time"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// SSHKeyReadArgs represents the arguments required to read an SSH key by ID or Name.
// It contains the SSH key ID or Name that is needed to perform the lookup.
type SSHKeyReadArgs struct {
	IDOrName string `json:"id_or_name" jsonschema:"required,description=The ssh key id or name to be searched"`
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
		Name:        "get_a_ssh_key_by_id_or_name",
		Description: "Retrieves a SSH key by its ID or Name, Get retrieves a SSH key by its ID if the input can be parsed as an integer, otherwise it retrieves a SSH key by its name. If the SSH key does not exist, nil is returned.",
		Handler: func(args SSHKeyReadArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*SSHKeyResponse, error) {
				result, _, err := client.SSHKey.Get(context.Background(), args.IDOrName)
				return tossSSHKeyResponse(result), err
			})
		},
		Restriction: RestrictionReadOnly,
	},
}
