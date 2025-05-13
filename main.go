// Package main is a Go Model Context Protocol (MCP) server for interacting with the Hetzner Cloud API.
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	"github.com/joho/godotenv"
	mcpgolang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

var client *hcloud.Client

// Define a constant for the empty string
const emptyString = ""

// NoArgs represents an empty structure used when no arguments are required.
type NoArgs struct{}

// SSHKeyReadArgs represents the arguments required to read an SSH key.
// It contains the SSH key ID that is needed to perform the lookup.
type SSHKeyReadArgs struct {
	SSHKeyID int64 `json:"ssh_key_id" jsonschema:"required,description=The ssh-key id to be searched"`
}

// SSHKeyCreateArgs represents the arguments required to create a new SSH key.
// It includes the name of the key, the public key itself, and optional user-defined labels.
type SSHKeyCreateArgs struct {
	Name      string            `json:"name" jsonschema:"required,description:Name of the SSH key"`
	PublicKey string            `json:"public_key" jsonschema:"required,description:Public key"`
	Labels    map[string]string `json:"labels,omitempty" jsonschema:"description:Optional, User-defined labels (key/value pairs) for the Resource."`
}

// SSHKeyUpdateArgs represents the arguments required to update an existing SSH key.
// It includes the SSH key ID, the new name of the SSH key, and optional user-defined labels.
type SSHKeyUpdateArgs struct {
	SSHKeyID int64             `json:"ssh_key_id" jsonschema:"required,description=The ssh-key id to be searched"`
	Name     string            `json:"name" jsonschema:"required,description:Name of the SSH key"`
	Labels   map[string]string `json:"labels,omitempty" jsonschema:"description:Optional, User-defined labels (key/value pairs) for the Resource."`
}

// LocationReadArgs represents the arguments required to read a location.
// It contains the location ID that is needed to perform the lookup.
type LocationReadArgs struct {
	LocationID int64 `json:"location_id" jsonschema:"required,description=The location id to be searched"`
}

// DatacenterReadArgs represents the arguments required to read a datacenter.
// It contains the datacenter ID that is needed to perform the lookup.
type DatacenterReadArgs struct {
	DatacenterID int64 `json:"datacenter_id" jsonschema:"required,description=The datacenter id to be searched"`
}

// ListArgs represents the arguments for listing resources.
// It includes pagination information like page number, items per page, and a label selector for filtering.
type ListArgs struct {
	Page          int    `json:"page" jsonschema:"description=Page (starting at 1)"`
	PerPage       int    `json:"per_page" jsonschema:"description=Items per page (0 means default)"`
	LabelSelector string `json:"label_selector" jsonschema:"description=Label selector for filtering by labels"`
}

// Tool represents a tool with a name, description, and handler function.
type Tool struct {
	Name        string
	Description string
	Handler     any
}

// Tools
var tools = []Tool{
	// SSH Keys
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
		Name:        "create_a_ssh_key",
		Description: "Creates a new SSH key with the given name and public_key. Once an SSH key is created, it can be used in other calls such as creating Servers.",
		Handler: func(args SSHKeyCreateArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.SSHKey, error) {
				result, _, err := client.SSHKey.Create(context.Background(), hcloud.SSHKeyCreateOpts{Name: args.Name, PublicKey: args.PublicKey, Labels: args.Labels})
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
	{
		Name:        "update_a_ssh_key",
		Description: "Updates a SSH key by its ID. You can update an SSH key name and an SSH key labels.",
		Handler: func(args SSHKeyUpdateArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.SSHKey, error) {
				sshKey, _, err := client.SSHKey.GetByID(context.Background(), args.SSHKeyID)
				if err != nil {
					return nil, err
				}
				result, _, err := client.SSHKey.Update(context.Background(), sshKey, hcloud.SSHKeyUpdateOpts{Name: args.Name, Labels: args.Labels})
				return result, err
			})
		},
	},
	{
		Name:        "delete_ssh_key_by_id",
		Description: "Deletes permanently a SSH key by its ID",
		Handler: func(args SSHKeyUpdateArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.SSHKey, error) {
				sshKey, _, err := client.SSHKey.GetByID(context.Background(), args.SSHKeyID)
				if err != nil {
					return nil, err
				}
				_, err = client.SSHKey.Delete(context.Background(), sshKey)
				return sshKey, err
			})
		},
	},

	// Location
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

	// Datacenter
	{
		Name:        "get_datacenter_list",
		Description: "Returns all datacenters objects",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.Datacenter, error) {
				result, _, err := client.Datacenter.List(context.Background(), hcloud.DatacenterListOpts{})
				return result, err
			})
		},
	},
	{
		Name:        "get_datacenter_info_by_id",
		Description: "Get a datacenter by its ID, it returns the datacenter object info",
		Handler: func(args DatacenterReadArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Datacenter, error) {
				result, _, err := client.Datacenter.GetByID(context.Background(), args.DatacenterID)
				return result, err
			})
		},
	},
}

// Generalized response handler for listing and getting server/location info
func handleResponse[T any](fetchFunc func() (T, error)) (*mcpgolang.ToolResponse, error) {
	// Fetch data using the provided fetch function
	data, err := fetchFunc()
	if err != nil {
		return nil, err
	}

	// Marshal the data into the desired format
	marshaledData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %w", err)
	}

	// Return the response
	return mcpgolang.NewToolResponse(mcpgolang.NewTextContent(string(marshaledData))), nil
}

// LoadToken loads the Hetzner Cloud API token from the command-line flag, environment variable, or .env file
func loadToken() string {
	// Define command-line flag
	tokenFlag := flag.String("token", "", "Hetzner Cloud API token")
	flag.Parse()

	// Try to load .env file from executable directory
	exPath, err := os.Executable()
	if err == nil {
		envPath := filepath.Join(filepath.Dir(exPath), ".env")
		_ = godotenv.Load(envPath) // ignore error
	}

	// Get token from command-line flag
	hcloudToken := *tokenFlag

	// If flag not provided, check environment variable
	if hcloudToken == emptyString {
		hcloudToken = os.Getenv("HCLOUD_TOKEN")
	}

	// If still empty, show error and help
	if hcloudToken == emptyString {
		fmt.Println("Error: HCLOUD_TOKEN is not set.\n" +
			"You can provide the token using one of the following methods:\n" +
			"1. Command line: ./mcp-hetzner -token=your_token_here\n" +
			"2. Environment var: export HCLOUD_TOKEN=your_token_here && ./mcp-hetzner\n" +
			"3. .env file: create a .env file with HCLOUD_TOKEN=your_token_here")
	}

	//nolint:unhandled-error
	fmt.Println("Token loaded successfully.")
	return hcloudToken
}

// Register Tools
func registerTools(server *mcpgolang.Server) error {
	all := [][]Tool{
		tools,
		firewallTools,
		serverTools,
	}

	var allTools []Tool
	for _, group := range all {
		allTools = append(allTools, group...)
	}

	for _, tool := range allTools {
		if err := server.RegisterTool(tool.Name, tool.Description, tool.Handler); err != nil {
			return fmt.Errorf("failed to register tool %s: %w", tool.Name, err)
		}
	}
	return nil
}

// Start
func main() {
	done := make(chan struct{})

	// New Stdio Server
	server := mcpgolang.NewServer(stdio.NewStdioServerTransport())

	// Load Hetzner Cloud token
	hcloudToken := loadToken()

	// Hetzner Cloud Client
	client = hcloud.NewClient(hcloud.WithToken(hcloudToken))

	// Register Tool
	err := registerTools(server)
	if err != nil {
		panic(err)
	}

	// Run server
	err = server.Serve()
	if err != nil {
		panic(err)
	}

	<-done
}
