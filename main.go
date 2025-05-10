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
	mcp_golang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

var client *hcloud.Client

type NoArgs struct{}

type SSHKeyReadArgs struct {
	SshKeyId int64 `json:"ssh_key_id" jsonschema:"required,description=The ssh-key id to be searched"`
}

type SSHKeyCreateArgs struct {
	Name      string            `json:"name" jsonschema:"required,description:Name of the SSH key"`
	PublicKey string            `json:"public_key" jsonschema:"required,description:Public key"`
	Labels    map[string]string `json:"labels,omitempty" jsonschema:"description:Optional, User-defined labels (key/value pairs) for the Resource."`
}

type SSHKeyUpdateArgs struct {
	SshKeyId int64             `json:"ssh_key_id" jsonschema:"required,description=The ssh-key id to be searched"`
	Name     string            `json:"name" jsonschema:"required,description:Name of the SSH key"`
	Labels   map[string]string `json:"labels,omitempty" jsonschema:"description:Optional, User-defined labels (key/value pairs) for the Resource."`
}

type ServerReadArgs struct {
	ServerId int64 `json:"server_id" jsonschema:"required,description=The server id to be searched"`
}

type LocationReadArgs struct {
	LocationId int64 `json:"location_id" jsonschema:"required,description=The location id to be searched"`
}

type Tool struct {
	Name        string
	Description string
	Handler     interface{}
}

// Tools
var tools = []Tool{
	// SSH Keys
	{
		Name:        "get_ssh_key_list",
		Description: "Returns all ssh-key objects. SSH keys are public keys you provide to the cloud system. They can be injected into Servers at creation time.",
		Handler: func(args NoArgs) (*mcp_golang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.SSHKey, error) {
				result, _, err := client.SSHKey.List(context.Background(), hcloud.SSHKeyListOpts{})
				return result, err
			})
		},
	},
	{
		Name:        "create_a_ssh_key",
		Description: "Creates a new SSH key with the given name and public_key. Once an SSH key is created, it can be used in other calls such as creating Servers.",
		Handler: func(args SSHKeyCreateArgs) (*mcp_golang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.SSHKey, error) {
				result, _, err := client.SSHKey.Create(context.Background(), hcloud.SSHKeyCreateOpts{Name: args.Name, PublicKey: args.PublicKey, Labels: args.Labels})
				return result, err
			})
		},
	},
	{
		Name:        "get_ssh_key_by_id",
		Description: "Get a SSH key by its ID, it returns a specific ssh key object info",
		Handler: func(args SSHKeyReadArgs) (*mcp_golang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.SSHKey, error) {
				result, _, err := client.SSHKey.GetByID(context.Background(), args.SshKeyId)
				return result, err
			})
		},
	},
	{
		Name:        "update_a_ssh_key",
		Description: "Updates a SSH key by its ID. You can update an SSH key name and an SSH key labels.",
		Handler: func(args SSHKeyUpdateArgs) (*mcp_golang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.SSHKey, error) {
				ssh_key, _, err := client.SSHKey.GetByID(context.Background(), args.SshKeyId)
				if err != nil {
					return nil, err
				}
				result, _, err := client.SSHKey.Update(context.Background(), ssh_key, hcloud.SSHKeyUpdateOpts{Name: args.Name, Labels: args.Labels})
				return result, err
			})
		},
	},
	{
		Name:        "delete_ssh_key_by_id",
		Description: "Deletes permanently a SSH key by its ID",
		Handler: func(args SSHKeyUpdateArgs) (*mcp_golang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.SSHKey, error) {
				ssh_key, _, err := client.SSHKey.GetByID(context.Background(), args.SshKeyId)
				if err != nil {
					return nil, err
				}
				_, err = client.SSHKey.Delete(context.Background(), ssh_key)
				return ssh_key, err
			})
		},
	},

	// Location
	{
		Name:        "get_location_list",
		Description: "Returns all locations objects",
		Handler: func(args NoArgs) (*mcp_golang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.Location, error) {
				result, _, err := client.Location.List(context.Background(), hcloud.LocationListOpts{})
				return result, err
			})
		},
	},
	{
		Name:        "get_location_info_by_id",
		Description: "Get a location by its ID, it returns the location object info",
		Handler: func(args LocationReadArgs) (*mcp_golang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Location, error) {
				result, _, err := client.Location.GetByID(context.Background(), args.LocationId)
				return result, err
			})
		},
	},

	// Server
	{
		Name:        "get_server_list",
		Description: "Returns all existing Server objects",
		Handler: func(args NoArgs) (*mcp_golang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.Server, error) {
				result, _, err := client.Server.List(context.Background(), hcloud.ServerListOpts{})
				return result, err
			})
		},
	},
	{
		Name:        "get_server_info_by_id",
		Description: "Get a server by its ID, it returns the server object info",
		Handler: func(args ServerReadArgs) (*mcp_golang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Server, error) {
				result, _, err := client.Server.GetByID(context.Background(), args.ServerId)
				return result, err
			})
		},
	},
}

// Generalized response handler for listing and getting server/location info
func handleResponse[T any](fetchFunc func() (T, error)) (*mcp_golang.ToolResponse, error) {
	// Fetch data using the provided fetch function
	data, err := fetchFunc()
	if err != nil {
		return nil, err
	}

	// Marshal the data into the desired format
	marshaled_data, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %w", err)
	}

	// Return the response
	return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(string(marshaled_data))), nil
}

// LoadToken loads the Hetzner Cloud API token from the command-line flag, environment variable, or .env file
func loadToken() string {
	// Define command-line flag
	token_flag := flag.String("token", "", "Hetzner Cloud API token")
	flag.Parse()

	// Try to load .env file from executable directory
	ex_path, err := os.Executable()
	if err == nil {
		env_path := filepath.Join(filepath.Dir(ex_path), ".env")
		_ = godotenv.Load(env_path) // ignore error
	}

	// Get token from command-line flag
	hcloud_token := *token_flag

	// If flag not provided, check environment variable
	if hcloud_token == "" {
		hcloud_token = os.Getenv("HCLOUD_TOKEN")
	}

	// If still empty, show error and help
	if hcloud_token == "" {
		fmt.Println("Error: HCLOUD_TOKEN is not set.")
		fmt.Println("You can provide the token using one of the following methods:")
		fmt.Println("1. Command line: ./mcp-hetzner -token=your_token_here")
		fmt.Println("2. Environment var: export HCLOUD_TOKEN=your_token_here && ./mcp-hetzner")
		fmt.Println("3. .env file: create a .env file with HCLOUD_TOKEN=your_token_here")
		os.Exit(1)
	}

	fmt.Println("Token loaded successfully.")
	return hcloud_token
}

// Register Tools
func registerTools(server *mcp_golang.Server) error {
	for _, tool := range tools {
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
	server := mcp_golang.NewServer(stdio.NewStdioServerTransport())

	// Load Hetzner Cloud token
	hcloud_token := loadToken()

	// Hetzner Cloud Client
	client = hcloud.NewClient(hcloud.WithToken(hcloud_token))

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
