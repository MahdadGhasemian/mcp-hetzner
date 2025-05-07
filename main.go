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

type Arguments struct {
	ServerId int64 `json:"server_id" jsonschema:"required,description=The server id to be searched"`
}

type ToolHandler func(arguments Arguments) (*mcp_golang.ToolResponse, error)

type Tool struct {
	Name        string
	Description string
	Handler     ToolHandler
}

var tools = []Tool{
	//
	{
		Name:        "get_server_list",
		Description: "Returns all existing Server objects",
		Handler: func(_ Arguments) (*mcp_golang.ToolResponse, error) {
			servers, _, err := client.Server.List(context.Background(), hcloud.ServerListOpts{})
			if err != nil {
				return nil, err
			}
			ids := make([]int64, len(servers))
			for i, server := range servers {
				ids[i] = server.ID
			}
			return mcp_golang.NewToolResponse(
				mcp_golang.NewTextContent(fmt.Sprintf("The server list is %v!", ids)),
			), nil
		},
	},
	//
	{
		Name:        "get_server_info_by_id",
		Description: "Get a server by its ID, it returns the server name",
		Handler: func(args Arguments) (*mcp_golang.ToolResponse, error) {
			server, _, err := client.Server.GetByID(context.Background(), args.ServerId)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(
				mcp_golang.NewTextContent(fmt.Sprintf("The server name is %s!", server.Name)),
			), nil
		},
	},
	//
	{
		Name:        "get_location_list",
		Description: "Get a list of locations, it returns the location names",
		Handler: func(_ Arguments) (*mcp_golang.ToolResponse, error) {
			locations, _, err := client.Location.List(context.Background(), hcloud.LocationListOpts{})
			if err != nil {
				return nil, err
			}

			location_json, err := json.MarshalIndent(locations, "", "  ")
			if err != nil {
				return nil, err
			}

			return mcp_golang.NewToolResponse(
				mcp_golang.NewTextContent(string(location_json)),
			), nil
		},
	},
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
		err := server.RegisterTool(tool.Name, tool.Description, tool.Handler)
		if err != nil {
			return err
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

	err = server.Serve()
	if err != nil {
		panic(err)
	}

	<-done
}
