package resources

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/transprogrammer/xenia/generated/hashicorp/azurerm/networksecuritygroup"
	"github.com/transprogrammer/xenia/generated/hashicorp/azurerm/resourcegroup"
	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
	"github.com/transprogrammer/xenia/generated/naming"
	"honnef.co/go/tools/config"
)

func NewVNet(stk cdktf.TerraformStack, cfg config.Config, naming *naming.Naming, rg *resourcegroup.ResourceGroup, addrSpace *[]*string, subnetInputs []vnet.VirtualNetworkSubnet) *vnet.VirtualNetwork {
	input := &vnet.VirtualNetworkConfig{
		Name:              (*naming).VirtualNetworkOutput(),
		AddressSpace:      addrSpace,
		Location:          (*rg).Location(),
		ResourceGroupName: (*rg).Name(),
		Subnet:            subnetInputs,
	}

	vnet := vnet.NewVirtualNetwork(stk, Ids.VirtualNetwork, input)

	return &vnet
}

func NewSubnetInput(stk cdktf.TerraformStack, naming naming.Naming, nsg networksecuritygroup.NetworkSecurityGroup, addressPrefix string) vnet.VirtualNetworkSubnet {

	return vnet.VirtualNetworkSubnet{
		Name:          naming.SubnetOutput(),
		AddressPrefix: &addressPrefix,
		SecurityGroup: nsg.Id(),
	}
}

func GetSubnet(vnet vnet.VirtualNetwork, index int) vnet.VirtualNetworkSubnetOutputReference {
	i := float64(index)

	return vnet.Subnet().Get(&i)
}
