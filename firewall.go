package main

import (
	"context"
	"net"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// An IPNet represents an IP network.
type IPNet struct {
	IP   []byte `json:"ip" jsonschema:"required,description=Network Number"`
	Mask []byte `json:"mask" jsonschema:"required,description=Network Mask"`
}

// FirewallRule represents a Firewall's rules.
type FirewallRule struct {
	Direction      string  `json:"direction" jsonschema:"required"`
	SourceIPs      []IPNet `json:"source_ips,omitempty"`
	DestinationIPs []IPNet `json:"destination_ips,omitempty"`
	Protocol       string  `json:"protocol" jsonschema:"required"`
	Port           *string `json:"port,omitempty"`
	Description    *string `json:"description,omitempty"`
}

// FirewallResourceServer represents a Server to apply a Firewall on.
type FirewallResourceServer struct {
	ID int64 `json:"id" jsonschema:"description=Server ID"`
}

// FirewallResourceLabelSelector represents a LabelSelector to apply a Firewall on.
type FirewallResourceLabelSelector struct {
	Selector string `json:"selector" jsonschema:"description=Label selector"`
}

// FirewallResource represents a resource to apply the new Firewall on.
type FirewallResource struct {
	Type          string                         `json:"type"`
	Server        *FirewallResourceServer        `json:"server,omitempty"`
	LabelSelector *FirewallResourceLabelSelector `json:"label_selector,omitempty"`
}

// FirewallReadArgs represents the arguments required to fetch a specific firewall by its ID.
type FirewallReadArgs struct {
	FirewallID int64 `json:"firewall_id" jsonschema:"required,description=The firewall id to be searched"`
}

// FirewallReadListArgs defines the filtering and sorting options for listing firewalls.
type FirewallReadListArgs struct {
	ListArgs
	Name string   `json:"name" jsonschema:"description=Filter resources by their name."`
	Sort []string `json:"sort" jsonschema:"description=Sort fields"`
}

// FirewallCreateArgs contains the necessary fields to create a new firewall,
// including its name, labels, rule definitions, and resources to apply to.
type FirewallCreateArgs struct {
	Name    string             `json:"name" jsonschema:"required"`
	Labels  map[string]string  `json:"labels,omitempty"`
	Rules   []FirewallRule     `json:"rules"`
	ApplyTo []FirewallResource `json:"apply_to"`
}

// FirewallUpdateArgs defines the parameters for updating an existing firewall,
// including the firewall ID, new name, updated labels, and optionally updated rules.
type FirewallUpdateArgs struct {
	FirewallID int64             `json:"firewall_id" jsonschema:"required"`
	Name       string            `json:"name" jsonschema:"required"`
	Labels     map[string]string `json:"labels,omitempty"`
	Rules      []FirewallRule    `json:"rules,omitempty"`
}

// FirewallDeleteArgs embeds FirewallReadArgs and is used to delete a firewall
// based on its unique ID.
type FirewallDeleteArgs struct {
	FirewallReadArgs
}

func convertRules(rules []FirewallRule) []hcloud.FirewallRule {
	converted := make([]hcloud.FirewallRule, len(rules))
	for i, rule := range rules {
		var port *string
		if rule.Port != nil {
			port = rule.Port
		}

		var description *string
		if rule.Description != nil {
			description = rule.Description
		}

		converted[i] = hcloud.FirewallRule{
			Direction:      hcloud.FirewallRuleDirection(rule.Direction),
			SourceIPs:      convertIPNets(rule.SourceIPs),
			DestinationIPs: convertIPNets(rule.DestinationIPs),
			Protocol:       hcloud.FirewallRuleProtocol(rule.Protocol),
			Port:           port,
			Description:    description,
		}
	}
	return converted
}

func convertApplyTo(resources []FirewallResource) []hcloud.FirewallResource {
	const initialSliceCapacity = 0

	converted := make([]hcloud.FirewallResource, initialSliceCapacity, len(resources))
	for _, r := range resources {
		if res := convertResource(r); res != nil {
			converted = append(converted, *res)
		}
	}
	return converted
}

func convertResource(r FirewallResource) *hcloud.FirewallResource {
	switch r.Type {
	case "server":
		return convertServerResource(r.Server)
	case "label_selector":
		return convertLabelSelectorResource(r.LabelSelector)
	default:
		// Unknown type: skip or log
		return nil
	}
}

func convertServerResource(server *FirewallResourceServer) *hcloud.FirewallResource {
	if server == nil {
		return nil
	}
	return &hcloud.FirewallResource{
		Type: hcloud.FirewallResourceTypeServer,
		Server: &hcloud.FirewallResourceServer{
			ID: server.ID,
		},
	}
}

func convertLabelSelectorResource(selector *FirewallResourceLabelSelector) *hcloud.FirewallResource {
	if selector == nil {
		return nil
	}
	return &hcloud.FirewallResource{
		Type: hcloud.FirewallResourceTypeLabelSelector,
		LabelSelector: &hcloud.FirewallResourceLabelSelector{
			Selector: selector.Selector,
		},
	}
}

func convertIPNets(ipnets []IPNet) []net.IPNet {
	converted := make([]net.IPNet, len(ipnets))
	for i, ipnet := range ipnets {
		converted[i] = net.IPNet{
			IP:   net.IP(ipnet.IP),
			Mask: net.IPMask(ipnet.Mask),
		}
	}
	return converted
}

// FirewallTools
var firewallTools = []Tool{
	{
		Name:        "get_firewall_list",
		Description: "Returns all existing Firewall objects",
		Handler: func(args FirewallReadListArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.Firewall, error) {
				result, _, err := client.Firewall.List(context.Background(), hcloud.FirewallListOpts{
					ListOpts: hcloud.ListOpts(args.ListArgs),
					Name:     args.Name,
					Sort:     args.Sort,
				})
				return result, err
			})
		},
	},
	{
		Name:        "create_a_firewall",
		Description: "Creates a new Firewall.",
		Handler: func(args FirewallCreateArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (hcloud.FirewallCreateResult, error) {
				result, _, err := client.Firewall.Create(context.Background(), hcloud.FirewallCreateOpts{
					Name:    args.Name,
					Labels:  args.Labels,
					Rules:   convertRules(args.Rules),
					ApplyTo: convertApplyTo(args.ApplyTo),
				})
				return result, err
			})
		},
	},
	{
		Name:        "get_firewall_by_id",
		Description: "Get a Firewall by its ID, it returns a specific Firewall object info",
		Handler: func(args FirewallReadArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Firewall, error) {
				result, _, err := client.Firewall.GetByID(context.Background(), args.FirewallID)
				return result, err
			})
		},
	},
	{
		Name:        "update_a_firewall",
		Description: "Updates a Firewall by its ID. You can update a Firewall name and labels.",
		Handler: func(args FirewallUpdateArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Firewall, error) {
				firewall, _, err := client.Firewall.GetByID(context.Background(), args.FirewallID)
				if err != nil {
					return nil, err
				}
				result, _, err := client.Firewall.Update(context.Background(), firewall, hcloud.FirewallUpdateOpts{Name: args.Name, Labels: args.Labels})
				return result, err
			})
		},
	},
	{
		Name:        "delete_firewall_by_id",
		Description: "Deletes permanently a Firewall by its ID",
		Handler: func(args FirewallDeleteArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Firewall, error) {
				firewall, _, err := client.Firewall.GetByID(context.Background(), args.FirewallID)
				if err != nil {
					return nil, err
				}
				_, err = client.Firewall.Delete(context.Background(), firewall)
				return firewall, err
			})
		},
	},
}
