package stk

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/cfg"
	"github.com/transprogrammer/xenia/pkg/prv"
	"github.com/transprogrammer/xenia/pkg/res"

	cnf "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/dataazurermclientconfig"
	sn "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/subnet"
	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
)

type PostgresCoreBeat interface {
	CoreBeat
	VNet() vnet.VirtualNetwork
	Client() cnf.DataAzurermClientConfig
}

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

func NewPostgres(scope constructs.Construct, cfg cfg.Config, core PostgresCoreBeat, tokens []string) DefaultPostgresDrum {
	name := NewName(tokens)

	stk := cdktf.NewTerraformStack(scope, name)
	prv.NewAzureRM(stk)

	naming := core.Naming()
	subnet := core.Subnet()
	vnet := core.VNet()
	client := core.Client()

	rg := res.NewResourceGroup(stk, cfg, naming)

	zone := res.NewPrivateDNSZone(stk, rg)
	link := res.NewPrivateDNSZoneVNetLink(stk, cfg, naming, rg, zone, vnet)

	srv := res.NewPostgresFlexibleServer(stk, cfg, naming, rg, subnet, zone, link, client)

	return DefaultPostgresDrum{
		StackName_: name,
		Stack_:     stk,
	}
}
