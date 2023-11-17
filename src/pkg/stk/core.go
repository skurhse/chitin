package stk

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
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

var CoreSubnets = CoreSubnetsIndex{
	Jump:     jsii.String(jumpAddr),
	Postgres: jsii.String(postgresAddr),
}

var CoreSubnetIndices = CoreSubnetsIndicesIndex{
	Jump:     0,
	Postgres: 1,
}

func NewCore(scope constructs.Construct, cfg CoreConfig, tokens Tokens) DefaultCoreDrum {
	name := NewName(tokens.Core)

	stk := NewStack(scope, name)
	prv.NewAzureRM(stk)

	naming := mod.NewNaming(stk, tokens.Core)
	jumpNaming := mod.NewNaming(stk, tokens.Jump)
	postgresNaming := mod.NewNaming(stk, tokens.Postgres)

	rg := res.NewResourceGroup(stk, cfg, naming)
	client := res.NewDataAzurermClientConfig(stk)

	jumpASG := res.NewASG(stk, cfg, jumpNaming, rg)
	jumpSecurityRule := res.NewSSHSecurityRule(cfg.WhitelistIPs(), jumpASG)
	jumpNSG := res.NewNSG(stk, cfg, jumpNaming, rg, jumpSecurityRule)
	jumpSubnetInput := res.NewSubnetInput(stk, jumpNaming, jumpNSG, CoreSubnets.Jump)

	postgresAddrs := CoreSubnets.Postgres
	postgresSubnetInput := res.NewSubnetInput(stk, postgresNaming, nil, postgresAddrs)

	subnetInputs := make([]vnet.VirtualNetworkSubnet, 2)
	subnetInputs[CoreSubnetIndices.Jump] = jumpSubnetInput
	subnetInputs[CoreSubnetIndices.Postgres] = postgresSubnetInput

	vnet := res.NewVNet(stk, cfg, naming, rg, CoreAddrSpace, subnetInputs)

	jumpSubnet := res.GetSubnet(vnet, CoreSubnetIndices.Jump)
	postgresSubnet := res.GetSubnet(vnet, CoreSubnetIndices.Postgres)

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
