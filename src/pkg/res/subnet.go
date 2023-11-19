package res

import (
	"fmt"

	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	rg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/resourcegroup"
	sn "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/subnet"
	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
	nm "github.com/transprogrammer/xenia/generated/naming"
)

func NewSubnet(stk cdktf.TerraformStack, naming nm.Naming, rg rg.ResourceGroup, vnet vnet.VirtualNetwork, addrPrefix *string, token string) sn.Subnet {

	id := fmt.Sprintf("%s_%s", *Ids.Subnet, token)

	input := sn.SubnetConfig{
		Name:               naming.SubnetOutput(),
		ResourceGroupName:  rg.Name(),
		VirtualNetworkName: vnet.Name(),
		AddressPrefixes:    &[]*string{addrPrefix},
	}

	return sn.NewSubnet(stk, &id, &input)
}

func NewPostgresSubnetDelegation() sn.SubnetDelegation {
	return sn.SubnetDelegation{
		Name: jsii.String("fs"),
		ServiceDelegation: &sn.SubnetDelegationServiceDelegation{
			Name: jsii.String("Microsoft.DBforPostgreSQL/flexibleServers"),
			Actions: &[]*string{
				jsii.String("Microsoft.Network/virtualNetworks/subnets/join/action"),
			},
		},
	}
}

func NewDelegatedSubnet(stk cdktf.TerraformStack, naming nm.Naming, rg rg.ResourceGroup, vnet vnet.VirtualNetwork, delegation sn.SubnetDelegation, addrPrefix *string, token string) sn.Subnet {

	id := fmt.Sprintf("%s_%s", *Ids.Subnet, token)

	input := sn.SubnetConfig{
		Name:               naming.SubnetOutput(),
		ResourceGroupName:  rg.Name(),
		VirtualNetworkName: vnet.Name(),
		AddressPrefixes:    &[]*string{addrPrefix},
		Delegation:         delegation,
	}

	return sn.NewSubnet(stk, &id, &input)
}
