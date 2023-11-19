package stk

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/transprogrammer/xenia/pkg/cfg"
	"github.com/transprogrammer/xenia/pkg/prv"
	"github.com/transprogrammer/xenia/pkg/res"
)

func NewPostgres(scope constructs.Construct, cfg cfg.Config, core PostgresCoreBeat, tokens []string) DefaultPostgresDrum {
	name := NewStackName(tokens)

	stk := cdktf.NewTerraformStack(scope, name)
	prv.NewAzureRM(stk)

	naming := core.Naming()
	subnet := core.Subnet()
	vnet := core.VNet()
	client := core.Client()

	rg := res.NewResourceGroup(stk, cfg, naming)

	zone := res.NewPrivateDNSZone(stk, rg)
	link := res.NewPrivateDNSZoneVNetLink(stk, cfg, naming, rg, zone, vnet)

	res.NewPostgresFlexibleServer(stk, cfg, naming, rg, subnet, zone, link, client)

	return DefaultPostgresDrum{
		StackName_: name,
		Stack_:     stk,
	}
}
