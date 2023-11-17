package stk

import (
	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
	"github.com/transprogrammer/xenia/generated/naming"
)

type CoreBeat interface {
	Naming() naming.Naming
	Subnet() vnet.VirtualNetworkSubnetOutputReference
}
