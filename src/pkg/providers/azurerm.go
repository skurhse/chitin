package providers

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/transprogrammer/xenia/generated/hashicorp/azurerm/provider"
)

func NewAzureRM(scope constructs.Construct) provider.AzurermProvider {
	input := &provider.AzurermProviderConfig{
		Features: &provider.AzurermProviderFeatures{},
	}

	return provider.NewAzurermProvider(scope, ProviderIds.AzureRM, input)
}
