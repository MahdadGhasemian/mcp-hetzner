package main

import (
	"context"
	"net"
	"time"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// ServerReadByIDArgs represents the arguments required to read an Server by ID.
// It contains the Server ID that is needed to perform the lookup.
type ServerReadByIDArgs struct {
	ID int64 `json:"id" jsonschema:"required,description=The server id to be searched"`
}

// ServerReadByNameArgs represents the arguments required to read an Server by Name.
// It contains the Server Name that is needed to perform the lookup.
type ServerReadByNameArgs struct {
	Name string `json:"name" jsonschema:"required,description=The server name to be searched"`
}

type ServerPublicNet struct {
	IPv4 net.IP
	IPv6 net.IP
}

type ServerType struct {
	ID   int64
	Name string
}

type Datacenter struct {
	ID   int64
	Name string
}

type ServerProtection struct {
	Delete  bool
	Rebuild bool
}

type Volume struct {
	ID     int64
	Name   string
	Status string
}

type ServerResponse struct {
	ID              int64             `json:"id" jsonschema:"required,description=Unique identifier of the server"`
	Name            string            `json:"name" jsonschema:"required,description=The name of the server"`
	Status          string            `json:"status" jsonschema:"required,description=Current status of the server"`
	Created         time.Time         `json:"created" jsonschema:"required,description=Timestamp of when the server was created"`
	PublicNet       ServerPublicNet   `json:"public_net" jsonschema:"description=Public network IP addresses of the server"`
	ServerType      ServerType        `json:"server_type" jsonschema:"required,description=Type of the server"`
	Datacenter      Datacenter        `json:"datacenter" jsonschema:"description=Datacenter where the server is located"`
	IncludedTraffic uint64            `json:"included_traffic" jsonschema:"required,description=Amount of included traffic in bytes"`
	OutgoingTraffic uint64            `json:"outgoing_traffic" jsonschema:"required,description=Outgoing traffic in bytes"`
	IngoingTraffic  uint64            `json:"ingoing_traffic" jsonschema:"required,description=Ingoing traffic in bytes"`
	BackupWindow    string            `json:"backup_window" jsonschema:"required,description=Time window when backups occur"`
	RescueEnabled   bool              `json:"rescue_enabled" jsonschema:"required,description=Whether rescue mode is enabled"`
	Locked          bool              `json:"locked" jsonschema:"required,description=Whether the server is currently locked"`
	Protection      ServerProtection  `json:"protection" jsonschema:"required,description=Server protection settings"`
	Labels          map[string]string `json:"labels" jsonschema:"description=User-defined labels for the server"`
	Volumes         []Volume          `json:"volumes" jsonschema:"description=List of attached volumes"`
	PrimaryDiskSize int               `json:"primary_disk_size" jsonschema:"description=Size of the primary disk in GB"`
}

func toServerResponse(s *hcloud.Server) *ServerResponse {
	if s == nil {
		return nil
	}

	// Convert volumes from []*Volume to []Volume
	volumes := make([]Volume, 0, len(s.Volumes))
	for _, v := range s.Volumes {
		if v != nil {
			volumes = append(volumes, Volume{
				ID:     v.ID,
				Name:   v.Name,
				Status: string(v.Status),
			})
		}
	}

	return &ServerResponse{
		ID:      s.ID,
		Name:    s.Name,
		Status:  string(s.Status),
		Created: s.Created,
		PublicNet: ServerPublicNet{
			IPv4: s.PublicNet.IPv4.IP,
			IPv6: s.PublicNet.IPv6.IP,
		},
		ServerType: ServerType{
			ID:   s.ServerType.ID,
			Name: s.ServerType.Name,
		},
		Datacenter: Datacenter{
			ID:   s.Datacenter.ID,
			Name: s.Datacenter.Name,
		},
		IncludedTraffic: s.IncludedTraffic,
		OutgoingTraffic: s.OutgoingTraffic,
		IngoingTraffic:  s.IngoingTraffic,
		BackupWindow:    s.BackupWindow,
		RescueEnabled:   s.RescueEnabled,
		Locked:          s.Locked,
		Protection: ServerProtection{
			Delete:  s.Protection.Delete,
			Rebuild: s.Protection.Rebuild,
		},
		Labels:          s.Labels,
		Volumes:         volumes,
		PrimaryDiskSize: s.PrimaryDiskSize,
	}
}

// ServerTools
var serverTools = []Tool{
	{
		Name:        "get_all_servers",
		Description: "Returns all Servers objects.",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*ServerResponse, error) {
				result, err := client.Server.All(context.Background())
				if err != nil {
					return nil, err
				}
				var filtered []*ServerResponse
				for _, s := range result {
					filtered = append(filtered, toServerResponse(s))
				}
				return filtered, nil
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_server_by_id",
		Description: "Retrieves a Server by its ID. If the Server does not exist, nil is returned.",
		Handler: func(args ServerReadByIDArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*ServerResponse, error) {
				result, _, err := client.Server.GetByID(context.Background(), args.ID)
				if err != nil {
					return nil, err
				}
				return toServerResponse(result), nil
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_server_by_name",
		Description: "Retrieves a Server by its Name. If the Server does not exist, nil is returned.",
		Handler: func(args ServerReadByNameArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*ServerResponse, error) {
				result, _, err := client.Server.GetByName(context.Background(), args.Name)
				if err != nil {
					return nil, err
				}
				return toServerResponse(result), nil
			})
		},
		Restriction: RestrictionReadOnly,
	},
}
