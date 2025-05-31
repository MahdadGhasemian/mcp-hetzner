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
	cwd, err := os.Getwd()
	if err == nil {
		envPath := filepath.Join(cwd, ".env")
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

	return hcloudToken
}

// Is Allowed
func isAllowed(toolRestriction, globalRestriction Restriction) bool {
	if globalRestriction == RestrictionReadWrite {
		return true // global allows everything
	}
	if globalRestriction == RestrictionReadOnly && toolRestriction == RestrictionReadOnly {
		return true // global read-only allows only read-only tools
	}
	return false
}

// Allowed Tools
func collectAllowedTools(restriction Restriction) []Tool {
	toolGroups := [][]Tool{
		certificateTools,
		locationTools,
		datacenterTools,
		sshkeyTools,
		firewallTools,
		floatingIPTools,
		serverTools,
		imageTools,
		isoTools,
		placementGroupTools,
		primaryIPTools,
		serverTypeTools,
		loadBalancerTools,
		loadBalancerTypeTools,
		networkTools,
		volumeTools,
		priceTools,
	}

	var allowed []Tool

	for _, group := range toolGroups {
		for _, tool := range group {
			if isAllowed(tool.Restriction, restriction) {
				allowed = append(allowed, tool)
			}
		}
	}

	return allowed
}

// Register Tools
func registerTools(server *mcpgolang.Server, restriction Restriction) error {
	allTools := collectAllowedTools(restriction)

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

	// Define CLI flag for restriction
	restrictionFlag := flag.String("restriction", string(RestrictionReadOnly), "Restriction for the tool")
	flag.Parse()

	// Validate restriction flag
	var restriction Restriction
	if *restrictionFlag == string(RestrictionReadOnly) {
		restriction = RestrictionReadOnly
	} else if *restrictionFlag == string(RestrictionReadWrite) {
		restriction = RestrictionReadWrite
	} else {
		panic("invalid restriction")
	}

	// New Stdio Server
	server := mcpgolang.NewServer(stdio.NewStdioServerTransport())

	// Load Hetzner Cloud token
	hcloudToken := loadToken()

	// Hetzner Cloud Client
	client = hcloud.NewClient(hcloud.WithToken(hcloudToken))

	// Register Tool
	err := registerTools(server, restriction)
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
