// Package main is a Go Model Context Protocol (MCP) server for interacting with the Hetzner Cloud API.
package main

import (
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
	if hcloudToken == EmptyString {
		hcloudToken = os.Getenv("HCLOUD_TOKEN")
	}

	// If still empty, show error and help
	if hcloudToken == EmptyString {
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
		certificateTools,
		locationTools,
		datacenterTools,
		sshkeyTools,
		firewallTools,
		floatingIPTools,
		serverTools,
		imageTools,
		isoTools,
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
