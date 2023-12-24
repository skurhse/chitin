package resourcegroup

import (
	"github.com/aws/constructs-go/constructs/v10"
	rg "github.com/skurhse/chitin/generated/hashicorp/azurerm/resourcegroup"
	"github.com/skurhse/chitin/generated/naming"
	"github.com/skurhse/chitin/pkg/cfg"
)

func NewResourceGroup(scope constructs.Construct, config cfg.Config, naming naming.Naming) rg.ResourceGroup {

	input := &rg.ResourceGroupConfig{
		Name:     naming.ResourceGroupOutput(),
		Location: config.Regions().Primary(),
	}

	return rg.NewResourceGroup(scope, Ids.ResourceGroup, input)
}
