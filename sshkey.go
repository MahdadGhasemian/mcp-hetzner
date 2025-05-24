package main

import (
	"context"

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

// SSHKeyTools
var sshkeyTools = []Tool{
	{
		Name:        "get_all_ssh_keys",
		Description: "Returns all ssh-key objects. SSH keys are public keys you provide to the cloud system. They can be injected into Servers at creation time.",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.SSHKey, error) {
				result, err := client.SSHKey.All(context.Background())
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_ssh_key_by_id",
		Description: "Retrieves a SSH key by its ID, it returns a specific ssh key object info. If the SSH key does not exist, nil is returned.",
		Handler: func(args SSHKeyReadByIDArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.SSHKey, error) {
				result, _, err := client.SSHKey.GetByID(context.Background(), args.ID)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_ssh_key_by_name",
		Description: "Retrieves a SSH key by its Name. If the SSH key does not exist, nil is returned.",
		Handler: func(args SSHKeyReadByNameArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.SSHKey, error) {
				result, _, err := client.SSHKey.GetByName(context.Background(), args.Name)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_ssh_key_by_fingerprint",
		Description: "Retrieves a SSH key by its Fingerprint. If the SSH key does not exist, nil is returned.",
		Handler: func(args SSHKeyReadByFingerprintArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.SSHKey, error) {
				result, _, err := client.SSHKey.GetByName(context.Background(), args.Fingerprint)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
}
