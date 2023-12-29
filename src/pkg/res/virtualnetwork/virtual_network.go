package virtualnetwork

import (
	"fmt"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/skurhse/chitin/generated/hashicorp/azurerm/resourcegroup"
	vnet "github.com/skurhse/chitin/generated/hashicorp/azurerm/virtualnetwork"
	"github.com/skurhse/chitin/generated/naming"
	"github.com/skurhse/chitin/pkg/res"
)

func NewVirtualNetwork(stk cdktf.TerraformStack, naming naming.Naming, rg resourcegroup.ResourceGroup, addrSpace []*string, token string) vnet.VirtualNetwork {

	id := fmt.Sprintf("%s_%s", res.Ids.VirtualNetwork, token)

	input := vnet.VirtualNetworkConfig{
		Name:              naming.VirtualNetworkOutput(),
		AddressSpace:      &addrSpace,
		Location:          rg.Location(),
		ResourceGroupName: rg.Name(),
	}

	return vnet.NewVirtualNetwork(stk, &id, &input)
}
