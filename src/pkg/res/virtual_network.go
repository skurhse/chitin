package res

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/transprogrammer/xenia/generated/hashicorp/azurerm/resourcegroup"
	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/cfg"
)

func NewVNet(stk cdktf.TerraformStack, cfg cfg.Config, naming naming.Naming, rg resourcegroup.ResourceGroup, addrSpace []*string, subnetInputs []vnet.VirtualNetworkSubnet) vnet.VirtualNetwork {
	input := vnet.VirtualNetworkConfig{
		Name:              naming.VirtualNetworkOutput(),
		AddressSpace:      &addrSpace,
		Location:          rg.Location(),
		ResourceGroupName: rg.Name(),
		Subnet:            subnetInputs,
	}

	return vnet.NewVirtualNetwork(stk, Ids.VirtualNetwork, &input)
}
