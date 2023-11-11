package stk

import (
	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
	"github.com/transprogrammer/xenia/generated/naming"
)

type DefaultClusterCoreBeat struct {
	Naming_ naming.Naming
	Subnet_ vnet.VirtualNetworkSubnetOutputReference
}

func (c DefaultClusterCoreBeat) Naming() naming.Naming {
	return c.Naming_
}

func (c DefaultClusterCoreBeat) Subnet() vnet.VirtualNetworkSubnetOutputReference {
	return c.Subnet_
}
