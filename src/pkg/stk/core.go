package stk

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/cfg"
	"github.com/transprogrammer/xenia/pkg/mod"
	"github.com/transprogrammer/xenia/pkg/prv"
	"github.com/transprogrammer/xenia/pkg/res"
)

type CoreDrum interface {
	Drum
	JumpBeat() JumpCoreBeat
	PostgresBeats() PostgresCoreBeats
	ClusterBeat() ClusterCoreBeat
}

type DefaultCoreDrum struct {
	StackName_   *string
	Stack_       cdktf.TerraformStack
	JumpBeat_    DefaultJumpCoreBeat
	PostgresBeats_  DefaultPostgresCoreBeats
	ClusterBeat_ DefaultClusterCoreBeat
}

type CoreConfig interface {
	cfg.Config
	JumpConfig
	PostgresConfig
}

type CoreRegions interface {
	Primary() string
	Secondary() string
}

type DefaultCoreConfig struct {
	Tokens_ []string
	Regions DefaultCoreRegions
}

type DefaultCoreRegions struct {
	Primary_   string
	Secondary_ string
}

type CoreBeat interface {
	Naming() naming.Naming
	Subnet() vnet.VirtualNetworkSubnetOutputReference
}

func (c DefaultCoreDrum) StackName() *string {
	return c.StackName_
}

func (c DefaultCoreDrum) Stack() cdktf.TerraformStack {
	return c.Stack_
}

func (c DefaultCoreDrum) JumpBeat() JumpCoreBeat {
	return c.JumpBeat_
}

func (c DefaultCoreDrum) PostgresBeats() PostgresCoreBeats {
	return PostgresCoreBeats(c.PostgresBeats_)
}

func (c DefaultCoreDrum) ClusterBeat() ClusterCoreBeat {
	return ClusterCoreBeat(c.ClusterBeat_)
}

type CoreSubnetsIndex struct {
	Jump  *string
	Postgres PostgresSubnetsIndex
}

type PostgresSubnetsIndex struct {
	Dev  *string
	Prod *string
}

const (
	coreAddr      = "10.0.0.0/16"
	jumpAddr      = "10.1.0.0./24"
	postgresDevAddr  = "10.2.0.0./24"
	postgresProdAddr = "10.3.0.0./24"
)

var CoreAddrSpace = []*string{jsii.String(coreAddr)}

var CoreSubnets = CoreSubnetsIndex{
	Jump: jsii.String(jumpAddr),
	Postgres: PostgresSubnetsIndex{
		Dev:  jsii.String(postgresDevAddr),
		Prod: jsii.String(postgresProdAddr),
	},
}

type CoreSubnetsIndicesIndex struct {
	Jump  int
	Postgres PostgresCoreSubnetsIndicesIndex
}

type PostgresCoreSubnetsIndicesIndex struct {
	Dev  int
	Prod int
}

var CoreSubnetIndices = CoreSubnetsIndicesIndex{
	Jump: 0,
	Postgres: PostgresCoreSubnetsIndicesIndex{
		Dev:  1,
		Prod: 2,
	},
}

func NewCore(scope constructs.Construct, cfg CoreConfig, tokens Tokens) DefaultCoreDrum {
	name := NewName(tokens.Core)

	stack := NewStack(scope, name)
	prv.NewAzureRM(stack)

	postgresTokens := tokens.Postgres

	naming := mod.NewNaming(stack, tokens.Core)
	jumpNaming := mod.NewNaming(stack, tokens.Jump)
	postgresDevNaming := mod.NewNaming(stack, postgresTokens.Dev)
	postgresProdNaming := mod.NewNaming(stack, postgresTokens.Prod)

	rg := res.NewResourceGroup(stack, cfg, naming)

	jumpASG := res.NewASG(stack, cfg, jumpNaming, rg)

	jumpSecurityRule := res.NewSSHSecurityRule(cfg.WhitelistIPs(), jumpASG)

	jumpNSG := res.NewNSG(stack, cfg, jumpNaming, rg, jumpSecurityRule)

	jumpSubnetInput := res.NewSubnetInput(stack, jumpNaming, jumpNSG, CoreSubnets.Jump)

	postgresDevAddrs := CoreSubnets.Postgres.Dev
	postgresProdAddrs := CoreSubnets.Postgres.Prod

	postgresDevSubnetInput := res.NewSubnetInput(stack, postgresDevNaming, nil, postgresDevAddrs)
	postgresProdSubnetInput := res.NewSubnetInput(stack, postgresProdNaming, nil, postgresProdAddrs)

	subnetInputs := make([]vnet.VirtualNetworkSubnet, 3)
	subnetInputs[CoreSubnetIndices.Jump] = jumpSubnetInput
	subnetInputs[CoreSubnetIndices.Postgres.Dev] = postgresDevSubnetInput
	subnetInputs[CoreSubnetIndices.Postgres.Prod] = postgresProdSubnetInput

	vnet := res.NewVNet(stack, cfg, naming, rg, CoreAddrSpace, subnetInputs)

	jumpSubnet := res.GetSubnet(vnet, CoreSubnetIndices.Jump)
	postgresDevSubnet := res.GetSubnet(vnet, CoreSubnetIndices.Postgres.Dev)
	postgresProdSubnet := res.GetSubnet(vnet, CoreSubnetIndices.Postgres.Prod)

	return DefaultCoreDrum{
		StackName_: name,
		Stack_:     stack,
		JumpBeat_: DefaultJumpCoreBeat{
			Naming_: jumpNaming,
			Subnet_: jumpSubnet,
			ASG_:    jumpASG,
			NSG_:    jumpNSG,
		},
		PostgresBeats_: DefaultPostgresCoreBeats{
			Dev_: DefaultPostgresCoreBeat{
				Naming_: postgresDevNaming,
				Subnet_: postgresDevSubnet,
				VNet_:   vnet,
			},
			Prod_: DefaultPostgresCoreBeat{
				Naming_: postgresProdNaming,
				Subnet_: postgresProdSubnet,
				VNet_:   vnet,
			},
		},
	}
}
