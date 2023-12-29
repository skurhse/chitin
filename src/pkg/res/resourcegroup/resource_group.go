package resourcegroup

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	rg "github.com/skurhse/chitin/generated/hashicorp/azurerm/resourcegroup"
	"github.com/skurhse/chitin/generated/naming"
	"github.com/skurhse/chitin/pkg/cfg"
	"github.com/skurhse/chitin/pkg/res"
)

func NewResourceGroup(scope constructs.Construct, config cfg.Config, naming naming.Naming) rg.ResourceGroup {

	region := config.Regions().Primary()

	input := rg.ResourceGroupConfig{
		Name:     naming.ResourceGroupOutput(),
		Location: jsii.String(region),
	}

	name := fmt.Sprintf("%s_%s", res.Ids.ResourceGroup, region)

	return rg.NewResourceGroup(scope, &name, &input)
}
