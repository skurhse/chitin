package stk

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/transprogrammer/xenia/pkg/mod"
	"github.com/transprogrammer/xenia/pkg/prv"
	"github.com/transprogrammer/xenia/pkg/res"
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

func NewCore(scope constructs.Construct, cfg CoreConfig, tokenSets TokenSetsIndex, token string) DefaultCoreDrum {

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

	return DefaultCoreDrum{
		StackName_: name,
		Stack_:     stk,
		JumpBeat_: DefaultJumpCoreBeat{
			Naming_: jumpName,
			Subnet_: jumpSubnet,
			ASG_:    jumpASG,
			NSG_:    jumpNSG,
		},
		PostgresBeat_: DefaultPostgresCoreBeat{
			Naming_: pgName,
			Subnet_: pgSubnet,
			VNet_:   vnet,
			Client_: client,
		},
	}
}
