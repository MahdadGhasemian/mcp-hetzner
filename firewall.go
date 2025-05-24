package main

import (
	"context"
	"encoding/base64"
	"log"
	"net"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// An IPNet represents an IP network.
type IPNet struct {
	IP   string `json:"ip" jsonschema:"required,description=Network Number, base64 format, example: 0.0.0.0"`
	Mask string `json:"mask" jsonschema:"required,description=Network Mask, base64 format, example: AAAAAA=="`
}

// FirewallResourceServer represents a Server to apply a Firewall on.
type FirewallResourceServer struct {
	ID int64 `json:"id" jsonschema:"description=Server ID"`
}

// FirewallResourceLabelSelector represents a LabelSelector to apply a Firewall on.
type FirewallResourceLabelSelector struct {
	Selector string `json:"selector" jsonschema:"description=Label selector"`
}

// FirewallRule represents a Firewall's rules.
type FirewallRule struct {
	Direction      string  `json:"direction" jsonschema:"required"`
	SourceIPs      []IPNet `json:"source_ips"`
	DestinationIPs []IPNet `json:"destination_ips"`
	Protocol       string  `json:"protocol" jsonschema:"required"`
	Port           *string `json:"port,omitempty"`
	Description    *string `json:"description,omitempty"`
}

// FirewallResource represents a resource to apply the new Firewall on.
type FirewallResource struct {
	Type          string                        `json:"type"`
	Server        FirewallResourceServer        `json:"server,omitempty"`
	LabelSelector FirewallResourceLabelSelector `json:"label_selector,omitempty"`
}

// FirewallReadByIDArgs represents the arguments required to read an Firewall by ID.
// It contains the Firewall ID that is needed to perform the lookup.
type FirewallReadByIDArgs struct {
	ID int64 `json:"id" jsonschema:"required,description=The firewall id to be searched"`
}

// FirewallReadByNameArgs represents the arguments required to read an Firewall by Name.
// It contains the Firewall Name that is needed to perform the lookup.
type FirewallReadByNameArgs struct {
	Name string `json:"name" jsonschema:"required,description=The firewall name to be searched"`
}

// FirewallCreateArgs contains the necessary fields to create a new firewall,
// including its name, labels, rule definitions, and resources to apply to.
type FirewallCreateArgs struct {
	Name    string             `json:"name" jsonschema:"required,description=The firewall name"`
	Labels  map[string]string  `json:"labels,omitempty"`
	Rules   []FirewallRule     `json:"rules"`
	ApplyTo []FirewallResource `json:"apply_to"`
}

func convertIPNets(ipnets []IPNet) []net.IPNet {
	converted := make([]net.IPNet, 0, len(ipnets))

	for _, ipnet := range ipnets {
		ip := net.ParseIP(ipnet.IP)
		if ip == nil {
			log.Printf("Invalid IP: %s", ipnet.IP)
			continue
		}

		maskBytes, err := base64.StdEncoding.DecodeString(ipnet.Mask)
		if err != nil {
			log.Printf("Invalid Mask (base64 decode failed): %s", ipnet.Mask)
			continue
		}

		mask := net.IPMask(maskBytes)
		if len(mask) != net.IPv4len && len(mask) != net.IPv6len {
			log.Printf("Invalid Mask length: %d", len(mask))
			continue
		}

		converted = append(converted, net.IPNet{
			IP:   ip,
			Mask: mask,
		})
	}

	return converted
}

func convertRules(rules []FirewallRule) []hcloud.FirewallRule {
	converted := make([]hcloud.FirewallRule, len(rules))

	for i, rule := range rules {
		converted[i] = hcloud.FirewallRule{
			Direction:      hcloud.FirewallRuleDirection(rule.Direction),
			SourceIPs:      convertIPNets(rule.SourceIPs),
			DestinationIPs: convertIPNets(rule.DestinationIPs),
			Protocol:       hcloud.FirewallRuleProtocol(rule.Protocol),
			Port:           rule.Port,
			Description:    rule.Description,
		}
	}

	return converted
}

func convertServerResource(server FirewallResourceServer) *hcloud.FirewallResourceServer {
	firewallResourceServer := hcloud.FirewallResourceServer{
		ID: server.ID,
	}

	return &firewallResourceServer
}

func convertLabelSelector(selector FirewallResourceLabelSelector) *hcloud.FirewallResourceLabelSelector {
	firewallResourceServer := hcloud.FirewallResourceLabelSelector{
		Selector: selector.Selector,
	}

	return &firewallResourceServer
}

func convertApplyTo(resources []FirewallResource) []hcloud.FirewallResource {
	converted := make([]hcloud.FirewallResource, len(resources))

	for i, resource := range resources {
		converted[i] = hcloud.FirewallResource{
			Type:          hcloud.FirewallResourceType(resource.Type),
			Server:        convertServerResource(resource.Server),
			LabelSelector: convertLabelSelector(resource.LabelSelector),
		}
	}

	return converted
}

// FirewallTools
var firewallTools = []Tool{
	{
		Name:        "get_all_firewalls",
		Description: "Returns all Firewalls objects.",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() ([]*hcloud.Firewall, error) {
				result, err := client.Firewall.All(context.Background())
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_firewall_by_id",
		Description: "Retrieves a Firewall by its ID. If the Firewall does not exist, nil is returned.",
		Handler: func(args FirewallReadByIDArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Firewall, error) {
				result, _, err := client.Firewall.GetByID(context.Background(), args.ID)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "get_a_firewall_by_name",
		Description: "Retrieves a Firewall by its Name. If the Firewall does not exist, nil is returned.",
		Handler: func(args FirewallReadByNameArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (*hcloud.Firewall, error) {
				result, _, err := client.Firewall.GetByName(context.Background(), args.Name)
				return result, err
			})
		},
		Restriction: RestrictionReadOnly,
	},
	{
		Name:        "create_a_firewall",
		Description: "Create a new Firewall",
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
		Restriction: RestrictionReadWrite,
	},
}
