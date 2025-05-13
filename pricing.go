package main

import (
	"context"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	mcpgolang "github.com/metoro-io/mcp-golang"
)

// PriceTools
var priceTools = []Tool{
	{
		Name:        "get_pricing_information",
		Description: "Get retrieves pricing information.",
		Handler: func(_ NoArgs) (*mcpgolang.ToolResponse, error) {
			return handleResponse(func() (hcloud.Pricing, error) {
				result, _, err := client.Pricing.Get(context.Background())
				return result, err
			})
		},
	},
}
