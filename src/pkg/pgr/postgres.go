package sng

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/skurhse/xen/pkg/cfg"
	"github.com/skurhse/xen/pkg/mod"
	"github.com/skurhse/xen/pkg/prv"
	"github.com/skurhse/xen/pkg/res"
)

func NewPostgres(scope constructs.Construct, cfg cfg.Config, core PostgresCoreTune, tokens []string) DefaultPostgresMelody {
	name := NewStackName(tokens)

	stk := cdktf.NewTerraformStack(scope, name)
	prv.NewAzureRM(stk)

	naming := core.Naming()
	vnet := core.VirtualNetwork()
	client := core.Client()

	rg := res.NewResourceGroup(stk, cfg, naming)

	pgName := mod.NewNaming(stk, tokens)
	pgDelegation := res.NewPostgresSubnetDelegation()
	pgSubnet := res.NewDelegatedSubnet(stk, pgName, rg, vnet, pgDelegation, CoreSubnetAddrs.Postgres, Tokens.Postgres)

	zone := res.NewPrivateDNSZone(stk, rg)
	link := res.NewPrivateDNSZoneVNetLink(stk, cfg, naming, rg, zone, vnet)

	res.NewPostgresFlexibleServer(stk, cfg, naming, rg, pgSubnet, zone, link, client)

	return DefaultPostgresMelody{
		StackName_: name,
		Stack_:     stk,
	}
}
