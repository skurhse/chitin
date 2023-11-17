package stk

import (
	asg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/applicationsecuritygroup"
	nsg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/networksecuritygroup"
	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
	"github.com/transprogrammer/xenia/generated/naming"
)

type DefaultJumpCoreBeat struct {
	Naming_ naming.Naming
	Subnet_ vnet.VirtualNetworkSubnetOutputReference
	ASG_    asg.ApplicationSecurityGroup
	NSG_    nsg.NetworkSecurityGroup
	VNet_   vnet.VirtualNetwork
}

func (c DefaultJumpCoreBeat) Naming() naming.Naming {
	return c.Naming_
}

func (c DefaultJumpCoreBeat) Subnet() vnet.VirtualNetworkSubnetOutputReference {
	return c.Subnet_
}

func (c DefaultJumpCoreBeat) ASG() asg.ApplicationSecurityGroup {
	return c.ASG_
}

func (c DefaultJumpCoreBeat) NSG() nsg.NetworkSecurityGroup {
	return c.NSG_
}

func (c DefaultJumpCoreBeat) VNet() vnet.VirtualNetwork {
	return c.VNet_
}
