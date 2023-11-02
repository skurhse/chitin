package stk

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/cfg"
	"github.com/transprogrammer/xenia/pkg/prv"
	"github.com/transprogrammer/xenia/pkg/res"

	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
)

type PostgresDrum interface {
	Drum
}

type DefaultPostgresDrum struct {
	StackName_ *string
	Stack_     cdktf.TerraformStack
}

func (self DefaultPostgresDrum) StackName() *string {
	return self.StackName_
}

func (self DefaultPostgresDrum) Stack() cdktf.TerraformStack {
	return self.Stack_
}

type PostgresConfig interface {
	cfg.Config
}

type PostgresCoreBeats interface {
	Dev() PostgresCoreBeat
	Prod() PostgresCoreBeat
}

type PostgresCoreBeat interface {
	CoreBeat
	VNet() vnet.VirtualNetwork
}

type DefaultPostgresCoreBeats struct {
	Dev_  DefaultPostgresCoreBeat
	Prod_ DefaultPostgresCoreBeat
}

type DefaultPostgresCoreBeat struct {
	Naming_ naming.Naming
	Subnet_ vnet.VirtualNetworkSubnetOutputReference
	VNet_   vnet.VirtualNetwork
}

func (c DefaultPostgresCoreBeats) Dev() PostgresCoreBeat {
	return c.Dev_
}

func (c DefaultPostgresCoreBeats) Prod() PostgresCoreBeat {
	return c.Prod_
}

func (c DefaultPostgresCoreBeat) Naming() naming.Naming {
	return c.Naming_
}

func (c DefaultPostgresCoreBeat) Subnet() vnet.VirtualNetworkSubnetOutputReference {
	return c.Subnet_
}

func (c DefaultPostgresCoreBeat) VNet() vnet.VirtualNetwork {
	return c.VNet_
}

func NewPostgres(scope constructs.Construct, cfg cfg.Config, core PostgresCoreBeat, tokens []string) DefaultPostgresDrum {
	name := NewName(tokens)

	stk := cdktf.NewTerraformStack(scope, name)
	prv.NewAzureRM(stk)

	naming := core.Naming()
	subnet := core.Subnet()
	vnet := core.VNet()

	rg := res.NewResourceGroup(stk, cfg, naming)

	server := NewPostgresServer(stk, cfg, naming, rg)

	res.NewCosmosDBPostgresDatabase(stk, cfg, naming, rg, acct)

	zone := res.NewPrivateDNSZone(stk, rg)
	res.NewPrivateDNSZoneVNetLink(stk, cfg, naming, rg, zone, vnet)

	res.NewPrivateEndpoint(stk, cfg, naming, rg, acct, subnet, zone)

	return DefaultPostgresDrum{
		StackName_: name,
		Stack_:     stk,
	}
}
