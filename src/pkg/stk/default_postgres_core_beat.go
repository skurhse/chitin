package stk

import (
	cnf "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/dataazurermclientconfig"
	sn "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/subnet"
	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
	"github.com/transprogrammer/xenia/generated/naming"
)

type DefaultPostgresCoreBeat struct {
	Naming_ naming.Naming
	Subnet_ sn.Subnet
	VNet_   vnet.VirtualNetwork
	Client_ cnf.DataAzurermClientConfig
}

func (c DefaultPostgresCoreBeat) Naming() naming.Naming {
	return c.Naming_
}

func (c DefaultPostgresCoreBeat) Subnet() sn.Subnet {
	return c.Subnet_
}

func (c DefaultPostgresCoreBeat) VNet() vnet.VirtualNetwork {
	return c.VNet_
}

func (c DefaultPostgresCoreBeat) Client() cnf.DataAzurermClientConfig {
	return c.Client_
}
