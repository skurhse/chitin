package cre

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/skurhse/xen/pkg/mod"
	"github.com/skurhse/xen/pkg/prv"
	"github.com/skurhse/xen/pkg/res"
)

const (
	coreAddr = "10.0.0.0/16"
	jumpAddr = "10.1.0.0./24"
	pgAddr   = "10.2.0.0./24"
)

var CoreAddrSpace = []*string{jsii.String(coreAddr)}

var CoreSubnetAddrs = CoreSubnetAddrsIndex{
	Jump:     jsii.String(jumpAddr),
	Postgres: jsii.String(pgAddr),
}

func NewCore(scope constructs.Construct, cfg CoreConfig, tokenSets TokenSetsIndex, token string) DefaultCoreMelody {

	name := NewStackName(tokenSets.Core)
	stk := NewStack(scope, name)
	prv.NewAzureRM(stk)

	coreName := mod.NewNaming(stk, tokenSets.Core)
	rg := res.NewResourceGroup(stk, cfg, coreName)
	client := res.NewDataAzurermClientConfig(stk)
	vnet := res.NewVirtualNetwork(stk, coreName, rg, CoreAddrSpace, token)

	jumpName := mod.NewNaming(stk, tokenSets.Jump)
	jumpASG := res.NewASG(stk, cfg, jumpName, rg)
	jumpSecurityRule := res.NewSSHSecurityRule(cfg.WhitelistIPs(), jumpASG)
	jumpNSG := res.NewNSG(stk, cfg, jumpName, rg, jumpSecurityRule)
	jumpSubnet := res.NewSubnet(stk, jumpName, rg, vnet, CoreSubnetAddrs.Jump, Tokens.Jump)
	res.NewSubnetNSGAssoc(stk, jumpSubnet, jumpNSG, Tokens.Jump)

	pgName := mod.NewNaming(stk, tokenSets.Postgres)
	pgDelegation := res.NewPostgresSubnetDelegation()
	pgSubnet := res.NewDelegatedSubnet(stk, pgName, rg, vnet, pgDelegation, CoreSubnetAddrs.Postgres, Tokens.Postgres)

	return DefaultCoreMelody{
		StackName_: name,
		Stack_:     stk,
		JumpTune_: DefaultJumpCoreTune{
			Naming_: jumpName,
			Subnet_: jumpSubnet,
			ASG_:    jumpASG,
			NSG_:    jumpNSG,
		},
		PostgresTune_: DefaultPostgresCoreTune{
			Naming_: pgName,
			Subnet_: pgSubnet,
			VNet_:   vnet,
			Client_: client,
		},
	}
}
