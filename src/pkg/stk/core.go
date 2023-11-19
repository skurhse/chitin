package stk

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/transprogrammer/xenia/pkg/mod"
	"github.com/transprogrammer/xenia/pkg/prv"
	"github.com/transprogrammer/xenia/pkg/res"
)

const (
	coreAddr     = "10.0.0.0/16"
	jumpAddr     = "10.1.0.0./24"
	postgresAddr = "10.2.0.0./24"
)

var CoreAddrSpace = []*string{jsii.String(coreAddr)}

var CoreSubnetAddrs = CoreSubnetAddrsIndex{
	Jump:     jsii.String(jumpAddr),
	Postgres: jsii.String(postgresAddr),
}

func NewCore(scope constructs.Construct, cfg CoreConfig, tokenSets TokenSetsIndex) DefaultCoreDrum {
	name := NewName(tokenSets.Core)

	stk := NewStack(scope, name)
	prv.NewAzureRM(stk)

	naming := mod.NewNaming(stk, tokenSets.Core)
	jumpNaming := mod.NewNaming(stk, tokenSets.Jump)
	postgresNaming := mod.NewNaming(stk, tokenSets.Postgres)

	rg := res.NewResourceGroup(stk, cfg, naming)
	client := res.NewDataAzurermClientConfig(stk)

	jumpASG := res.NewASG(stk, cfg, jumpNaming, rg)
	jumpSecurityRule := res.NewSSHSecurityRule(cfg.WhitelistIPs(), jumpASG)
	jumpNSG := res.NewNSG(stk, cfg, jumpNaming, rg, jumpSecurityRule)
	jumpSubnet := res.NewSubnet(stk, jumpNaming, rg, vnet, jumpNSG, CoreSubnetAddrs.Jump, CoreSubnetAddrs.Jump, Tokens.Jump)

	postgresAddrs := CoreSubnets.Postgres
	postgresSubnet := res.NewSubnet(postgresNaming, nil, postgresAddrs)

	vnet := res.NewVNet(stk, cfg, naming, rg, CoreAddrSpace, subnetInputs)

	return DefaultCoreDrum{
		StackName_: name,
		Stack_:     stk,
		JumpBeat_: DefaultJumpCoreBeat{
			Naming_: jumpNaming,
			Subnet_: jumpSubnet,
			ASG_:    jumpASG,
			NSG_:    jumpNSG,
		},
		PostgresBeat_: DefaultPostgresCoreBeat{
			Naming_: postgresNaming,
			Subnet_: postgresSubnet,
			VNet_:   vnet,
			Client_: client,
		},
	}
}
