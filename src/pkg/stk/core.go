package stk

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/cfg"
	"github.com/transprogrammer/xenia/pkg/mod"
	"github.com/transprogrammer/xenia/pkg/providers"
	"github.com/transprogrammer/xenia/pkg/res"
)

type CoreDrum interface {
	Drum
	JumpBeat() JumpCoreBeat
	MongoBeats() MongoCoreBeats
	ClusterBeat() ClusterCoreBeat
}

type DefaultCoreDrum struct {
	StackName_   *string
	Stack_       cdktf.TerraformStack
	JumpBeat_    DefaultJumpCoreBeat
	MongoBeats_  DefaultMongoCoreBeats
	ClusterBeat_ DefaultClusterCoreBeat
}

type CoreConfig interface {
	cfg.Config
	JumpConfig
	MongoConfig
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

func (c DefaultCoreDrum) MongoBeats() MongoCoreBeats {
	return MongoCoreBeats(c.MongoBeats_)
}

func (c DefaultCoreDrum) ClusterBeat() ClusterCoreBeat {
	return ClusterCoreBeat(c.ClusterBeat_)
}

type CoreSubnetsIndex struct {
	Jump  *string
	Mongo MongoSubnetsIndex
}

type MongoSubnetsIndex struct {
	Dev  *string
	Prod *string
}

const (
	coreAddr      = "10.0.0.0/16"
	jumpAddr      = "10.1.0.0./24"
	mongoDevAddr  = "10.2.0.0./24"
	mongoProdAddr = "10.3.0.0./24"
)

var CoreAddrSpace = []*string{jsii.String(coreAddr)}

var CoreSubnets = CoreSubnetsIndex{
	Jump: jsii.String(jumpAddr),
	Mongo: MongoSubnetsIndex{
		Dev:  jsii.String(mongoDevAddr),
		Prod: jsii.String(mongoProdAddr),
	},
}

type CoreSubnetsIndicesIndex struct {
	Jump  int
	Mongo MongoCoreSubnetsIndicesIndex
}

type MongoCoreSubnetsIndicesIndex struct {
	Dev  int
	Prod int
}

var CoreSubnetIndices = CoreSubnetsIndicesIndex{
	Jump: 0,
	Mongo: MongoCoreSubnetsIndicesIndex{
		Dev:  1,
		Prod: 2,
	},
}

func NewCore(scope constructs.Construct, cfg CoreConfig, tokens Tokens) DefaultCoreDrum {
	name := NewName(tokens.Core)

	stack := NewStack(scope, name)
	providers.NewAzureRM(stack)

	mongoTokens := tokens.Mongo

	naming := mod.NewNaming(stack, tokens.Core)
	jumpNaming := mod.NewNaming(stack, tokens.Jump)
	mongoDevNaming := mod.NewNaming(stack, mongoTokens.Dev)
	mongoProdNaming := mod.NewNaming(stack, mongoTokens.Prod)

	rg := res.NewResourceGroup(stack, cfg, naming)

	jumpASG := res.NewASG(stack, cfg, jumpNaming, rg)

	jumpSecurityRule := res.NewSSHSecurityRule(cfg.WhitelistIPs(), jumpASG)

	jumpNSG := res.NewNSG(stack, cfg, jumpNaming, rg, jumpSecurityRule)

	jumpSubnetInput := res.NewSubnetInput(stack, jumpNaming, jumpNSG, CoreSubnets.Jump)

	mongoDevAddrs := CoreSubnets.Mongo.Dev
	mongoProdAddrs := CoreSubnets.Mongo.Prod

	mongoDevSubnetInput := res.NewSubnetInput(stack, mongoDevNaming, nil, mongoDevAddrs)
	mongoProdSubnetInput := res.NewSubnetInput(stack, mongoProdNaming, nil, mongoProdAddrs)

	subnetInputs := make([]vnet.VirtualNetworkSubnet, 3)
	subnetInputs[CoreSubnetIndices.Jump] = jumpSubnetInput
	subnetInputs[CoreSubnetIndices.Mongo.Dev] = mongoDevSubnetInput
	subnetInputs[CoreSubnetIndices.Mongo.Prod] = mongoProdSubnetInput

	vnet := res.NewVNet(stack, cfg, naming, rg, CoreAddrSpace, subnetInputs)

	jumpSubnet := res.GetSubnet(vnet, CoreSubnetIndices.Jump)
	mongoDevSubnet := res.GetSubnet(vnet, CoreSubnetIndices.Mongo.Dev)
	mongoProdSubnet := res.GetSubnet(vnet, CoreSubnetIndices.Mongo.Prod)

	return DefaultCoreDrum{
		StackName_: name,
		Stack_:     stack,
		JumpBeat_: DefaultJumpCoreBeat{
			Naming_: jumpNaming,
			Subnet_: jumpSubnet,
			ASG_:    jumpASG,
			NSG_:    jumpNSG,
		},
		MongoBeats_: DefaultMongoCoreBeats{
			Dev_: DefaultMongoCoreBeat{
				Naming_: mongoDevNaming,
				Subnet_: mongoDevSubnet,
				VNet_:   vnet,
			},
			Prod_: DefaultMongoCoreBeat{
				Naming_: mongoProdNaming,
				Subnet_: mongoProdSubnet,
				VNet_:   vnet,
			},
		},
	}
}
