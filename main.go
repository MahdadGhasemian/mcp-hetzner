package main

import (
	"context"
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

type MyFunctionsArguments struct {
	ServerId int64 `json:"server_id" jsonschema:"required,description=The server id to be searched"`
}

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

func main() {
	done := make(chan struct{})

	server := mcp_golang.NewServer(stdio.NewStdioServerTransport())

	// Load Hetzner Cloud token
	hcloud_token := loadToken()

	// Hetzner Cloud Client
	client = hcloud.NewClient(hcloud.WithToken(hcloud_token))

	// Register Tool
	err := server.RegisterTool("server", "Get a server by his id, it returns the server name", func(arguments MyFunctionsArguments) (*mcp_golang.ToolResponse, error) {
		serverName, err := getServerById(arguments.ServerId)
		if err != nil {
			return nil, err
		}

		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("The server name is %s!", serverName))), nil
	})
	if err != nil {
		panic(err)
	}

	// Register Tool
	err = server.RegisterTool("serverlist", "Get a list of servers, it returns the server id", func(arguments MyFunctionsArguments) (*mcp_golang.ToolResponse, error) {
		serverIds, err := getServerListId()
		if err != nil {
			return nil, err
		}

		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("The server list is %v!", serverIds))), nil
	})
	if err != nil {
		panic(err)
	}

	// Register Tool
	err = server.RegisterTool("locationlist", "Get a list of locations, it returns the location name", func(arguments MyFunctionsArguments) (*mcp_golang.ToolResponse, error) {
		locationNames, err := getLocationList()
		if err != nil {
			return nil, err
		}

		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("The location list is %v!", locationNames))), nil
	})
	if err != nil {
		panic(err)
	}

	err = server.Serve()
	if err != nil {
		panic(err)
	}

	<-done
}

func getServerById(id int64) (string, error) {
	server, _, err := client.Server.GetByID(context.Background(), id)
	if err != nil {
		return "", err
	}

	return server.Name, nil
}

func getServerListId() ([]int64, error) {
	servers, _, err := client.Server.List(context.Background(), hcloud.ServerListOpts{})
	if err != nil {
		return nil, err
	}

	serverIds := make([]int64, len(servers))
	for i, server := range servers {
		serverIds[i] = server.ID
	}

	return serverIds, nil
}

func getLocationList() ([]string, error) {
	locations, _, err := client.Location.List(context.Background(), hcloud.LocationListOpts{})
	if err != nil {
		return nil, err
	}

	locationNames := make([]string, len(locations))
	for i, location := range locations {
		locationNames[i] = location.Name
	}

	return locationNames, nil
}
