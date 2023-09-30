package resources

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	rg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/resourcegroup"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/apps"
)

func NewResourceGroup(stack cdktf.TerraformStack, cfg apps.Config, naming *naming.Naming) *rg.ResourceGroup {

	id := ResourceIds.ResourceGroup

	input := &rg.ResourceGroupConfig{
		Name:     (*naming).ResourceGroupOutput(),
		Location: cfg.Regions().Primary(),
	}

	resourceGroup := rg.NewResourceGroup(stack, id, input)

	return &resourceGroup
}
