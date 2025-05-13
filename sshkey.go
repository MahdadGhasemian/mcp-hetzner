package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// SSHKeyReadArgs represents the arguments required to read an SSH key.
// It contains the SSH key ID that is needed to perform the lookup.
type SSHKeyReadArgs struct {
	SSHKeyID int64 `json:"ssh_key_id" jsonschema:"required,description=The ssh-key id to be searched"`
}

// SSHKeyTools
var sshkeyTools = []Tool{
	{
		Name:        "get_ssh_key_list",
		Description: "Returns all ssh-key objects. SSH keys are public keys you provide to the cloud system. They can be injected into Servers at creation time.",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.SSHKey, error) {
				result, _, err := client.SSHKey.List(context.Background(), hcloud.SSHKeyListOpts{})
				return result, err
			})
		},
	},
	{
		Name:        "get_ssh_key_by_id",
		Description: "Get a SSH key by its ID, it returns a specific ssh key object info",
		Handler: func(args SSHKeyReadArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.SSHKey, error) {
				result, _, err := client.SSHKey.GetByID(context.Background(), args.SSHKeyID)
				return result, err
			})
		},
	},
}
