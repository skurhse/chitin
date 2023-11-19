package stk

import (
	asg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/applicationsecuritygroup"
	nsg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/networksecuritygroup"
	sn "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/subnet"
	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
	"github.com/transprogrammer/xenia/generated/naming"
)

type DefaultJumpCoreBeat struct {
	Naming_ naming.Naming
	Subnet_ sn.Subnet
	ASG_    asg.ApplicationSecurityGroup
	NSG_    nsg.NetworkSecurityGroup
	VNet_   vnet.VirtualNetwork
}

func (c DefaultJumpCoreBeat) Naming() naming.Naming {
	return c.Naming_
}

func (c DefaultJumpCoreBeat) Subnet() sn.Subnet {
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
