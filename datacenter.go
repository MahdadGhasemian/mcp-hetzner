package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// DatacenterReadByIDArgs represents the arguments required to read an Datacenter by ID.
// It contains the Datacenter ID that is needed to perform the lookup.
type DatacenterReadByIDArgs struct {
	ID int64 `json:"id" jsonschema:"required,description=The datacenter id to be searched"`
}

// DatacenterReadByNameArgs represents the arguments required to read an Datacenter by Name.
// It contains the Datacenter Name that is needed to perform the lookup.
type DatacenterReadByNameArgs struct {
	Name string `json:"name" jsonschema:"required,description=The datacenter name to be searched"`
}

type Location struct {
	ID   int64  `json:"id" jsonschema:"required,description=The location id"`
	Name string `json:"name" jsonschema:"required,description=The location name"`
}

type ServerTypeId struct {
	ID int64 `json:"id" jsonschema:"required,description=The server type id"`
}

type DatacenterResponse struct {
	ID                               int64          `json:"id" jsonschema:"required,description=The datacenter id"`
	Name                             string         `json:"name" jsonschema:"required,description=The datacenter name"`
	Description                      string         `json:"description" jsonschema:"description=The datacenter description"`
	Location                         Location       `json:"location" jsonschema:"required,description=The location of the datacenter"`
	SupportedServerTypes             []ServerTypeId `json:"supported_server_types" jsonschema:"description=List of supported server types"`
	AvailableForMigrationServerTypes []ServerTypeId `json:"available_for_migration_server_types" jsonschema:"description=List of available for migration server types"`
	AvailableServerTypes             []ServerTypeId `json:"available_server_types" jsonschema:"description=List of available server types"`
}

func toServerTypeIdList(serverTypes []*hcloud.ServerType) []ServerTypeId {
	var result []ServerTypeId
	for _, serverType := range serverTypes {
		result = append(result, ServerTypeId{ID: serverType.ID})
	}
	return result
}

func toDatacenterResponse(d *hcloud.Datacenter) *DatacenterResponse {
	if d == nil {
		return nil
	}

	return &DatacenterResponse{
		ID:          d.ID,
		Name:        d.Name,
		Description: d.Description,
		Location: Location{
			ID:   d.Location.ID,
			Name: d.Location.Name,
		},
		SupportedServerTypes:             toServerTypeIdList(d.ServerTypes.Supported),
		AvailableForMigrationServerTypes: toServerTypeIdList(d.ServerTypes.AvailableForMigration),
		AvailableServerTypes:             toServerTypeIdList(d.ServerTypes.Available),
	}
}

// DatacenterTools
var datacenterTools = []Tool{
	{
		Name:        "get_all_datacenters",
		Description: "Returns all Datacenters objects.",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*DatacenterResponse, error) {
				result, err := client.Datacenter.All(context.Background())
				if err != nil {
					return nil, err
				}
				var filtered []*DatacenterResponse
				for _, d := range result {
					filtered = append(filtered, toDatacenterResponse(d))
				}
				return filtered, nil
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_datacenter_by_id",
		Description: "Retrieves a Datacenter by its ID. If the Datacenter does not exist, nil is returned.",
		Handler: func(args DatacenterReadByIDArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*DatacenterResponse, error) {
				result, _, err := client.Datacenter.GetByID(context.Background(), args.ID)
				if err != nil {
					return nil, err
				}
				return toDatacenterResponse(result), nil
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_datacenter_by_name",
		Description: "Retrieves a Datacenter by its Name. If the Datacenter does not exist, nil is returned.",
		Handler: func(args DatacenterReadByNameArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*DatacenterResponse, error) {
				result, _, err := client.Datacenter.GetByName(context.Background(), args.Name)
				if err != nil {
					return nil, err
				}
				return toDatacenterResponse(result), nil
			})
		},
		Restriction: RestrictionReadOnly,
	},
}
