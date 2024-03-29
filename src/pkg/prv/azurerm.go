package prv

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/skurhse/chitin/generated/hashicorp/azurerm/provider"
)

func NewAzureRM(scope constructs.Construct) provider.AzurermProvider {
	input := &provider.AzurermProviderConfig{
		Features: &provider.AzurermProviderFeatures{},
	}

	return provider.NewAzurermProvider(scope, ProviderIds.AzureRM, input)
}
