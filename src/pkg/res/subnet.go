package res

import (
	"fmt"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	rg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/resourcegroup"
	sn "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/subnet"
	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
	nm "github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/cfg"
)

func NewSubnet(stk cdktf.TerraformStack, naming nm.Naming, rg rg.ResourceGroup, vnet vnet.VirtualNetwork, addrPrefix *string, token *string) sn.Subnet {

	id := fmt.Sprintf("%s_%s", Ids.Subnet, token)

	input := sn.SubnetConfig{
		Name:               naming.SubnetOutput(),
		ResourceGroupName:  rg.Name(),
		VirtualNetworkName: vnet.Name(),
		AddressPrefixes:    &[]*string{addrPrefix},
		Delegation: 
	}

	return sn.NewSubnet(stk, &id, &input)
}


func NewDelegation(
