package resourcegroup

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	rg "github.com/skurhse/chitin/generated/hashicorp/azurerm/resourcegroup"
	"github.com/skurhse/chitin/generated/naming"
	"github.com/skurhse/chitin/pkg/cfg"
	"github.com/skurhse/chitin/pkg/res"
)

func NewResourceGroup(scope constructs.Construct, config cfg.Config, naming naming.Naming) rg.ResourceGroup {

	input := &rg.ResourceGroupConfig{
		Name:     naming.ResourceGroupOutput(),
		Location: jsii.String(config.Regions().Primary()),
	}

	return rg.NewResourceGroup(scope, res.Ids.ResourceGroup, input)
}
