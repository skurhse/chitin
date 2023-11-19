package stk

import (
	cnf "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/dataazurermclientconfig"
	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
)

type PostgresCoreBeat interface {
	CoreBeat
	VNet() vnet.VirtualNetwork
	Client() cnf.DataAzurermClientConfig
}
