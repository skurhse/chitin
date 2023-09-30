package providers

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/transprogrammer/xenia/generated/hashicorp/azurerm/provider"
	"github.com/transprogrammer/xenia/pkg/apps"
)

func NewAzureRM(stack cdktf.TerraformStack, config apps.Config) *provider.AzurermProvider {
	id := ProviderIds.AzureRM

	input := &provider.AzurermProviderConfig{
		Features: &provider.AzurermProviderFeatures{},
	}

	provider := provider.NewAzurermProvider(stack, id, input)

	return &provider
}
