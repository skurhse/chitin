package res

import (
	"github.com/aws/constructs-go/constructs/v10"
	rg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/resourcegroup"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/cfg"
)

func NewResourceGroup(scope constructs.Construct, config cfg.Config, naming naming.Naming) rg.ResourceGroup {

	input := &rg.ResourceGroupConfig{
		Name:     naming.ResourceGroupOutput(),
		Location: config.Regions().Primary(),
	}

	return rg.NewResourceGroup(scope, Ids.ResourceGroup, input)
}
