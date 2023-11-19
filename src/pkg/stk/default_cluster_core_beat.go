package stk

import (
	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
	"github.com/transprogrammer/xenia/generated/naming"
)

type DefaultClusterCoreBeat struct {
	Naming_ naming.Naming
	Subnet_ sn.Subnet
}

func (c DefaultClusterCoreBeat) Naming() naming.Naming {
	return c.Naming_
}

func (c DefaultClusterCoreBeat) Subnet() sn.Subnet {
	return c.Subnet_
}
