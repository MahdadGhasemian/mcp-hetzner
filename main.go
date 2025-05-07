package main

import (
	"context"
	"fmt"
	"os"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	"github.com/joho/godotenv"
	mcp_golang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

var client *hcloud.Client

type MyFunctionsArguments struct {
	ServerId int64 `json:"server_id" jsonschema:"required,description=The server id to be searched"`
}

func main() {
	done := make(chan struct{})

	server := mcp_golang.NewServer(stdio.NewStdioServerTransport())

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	hcloud_token := os.Getenv("HCLOUD_TOKEN")

	// Hetzner Cloud Client
	client = hcloud.NewClient(hcloud.WithToken(hcloud_token))

	// Register Tool
	err = server.RegisterTool("server", "Get a server by his id, it returns the server name", func(arguments MyFunctionsArguments) (*mcp_golang.ToolResponse, error) {
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
