package stacks

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/cfg"
	"github.com/transprogrammer/xenia/pkg/modules"
	"github.com/transprogrammer/xenia/pkg/providers"
	"github.com/transprogrammer/xenia/pkg/resources"
)

type CoreDrum interface {
	Drum
	JumpBeat() JumpCoreBeat
	MongoBeats() MongoCoreBeats
}

type DefaultCoreDrum struct {
	StackName_  *string
	Stack_      cdktf.TerraformStack
	JumpBeat_   DefaultJumpCoreBeat
	MongoBeats_ DefaultMongoCoreBeats
}

type CoreConfig interface {
	cfg.Config
	JumpConfig
	MongoConfig
}

type CoreRegions interface {
	Primary() *string
	Secondary() *string
}

type DefaultCoreConfig struct {
	Tokens_ []string
	Regions DefaultCoreRegions
}

type DefaultCoreRegions struct {
	Primary_   *string
	Secondary_ *string
}

type CoreBeat interface {
	Naming() *naming.Naming
	Subnet() *vnet.VirtualNetworkSubnetOutputReference
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

var CoreAddressSpace = &[]*string{jsii.String("10.0.0.0/16")}

type CoreSubnetsIndex struct {
	Jump  *string
	Mongo MongoSubnetsIndex
}

type MongoSubnetsIndex struct {
	Development *string
	Production  *string
}

var CoreSubnets = CoreSubnetsIndex{
	Jump: jsii.String("10.1.0.0/24"),
	Mongo: MongoSubnetsIndex{
		Development: jsii.String("10.2.0.0/24"),
		Production:  jsii.String("10.3.0.0/24"),
	},
}

type CoreSubnetsIndicesIndex struct {
	Jump  int
	Mongo MongoCoreSubnetsIndicesIndex
}

type MongoCoreSubnetsIndicesIndex struct {
	Development int
	Production  int
}

var CoreSubnetIndices = CoreSubnetsIndicesIndex{
	Jump: 0,
	Mongo: MongoCoreSubnetsIndicesIndex{
		Development: 1,
		Production:  2,
	},
}

func NewCore(scope constructs.Construct, cfg CoreConfig, tokens ...string) DefaultCoreDrum {
	tokens = EnrichTokens(cfg)
	name := NewName(tokens)

	stack := NewStack(scope, name)

	stack := NewStack(scope, StackNames.Core)
	providers.NewAzureRM(stack)

	naming := modules.NewNaming(stack, StackTokens.Core)

	rg := resources.NewResourceGroup(stack, cfg, naming)

	jumpNaming := modules.NewNaming(stack, cfg, StackTokens.Jump)

	jumpASG := resources.NewAppSecurityGroup(stack, cfg, jumpNaming, rg)

	jumpSecurityRule := resources.NewSSHSecurityRule(cfg, jumpASG)

	jumpNSG := resources.NewNSG(stack, cfg, jumpNaming, rg, jumpSecurityRule)

	jumpSubnetInput := resources.NewSubnetInput(stack, jumpNaming, jumpNSG, CoreSubnets.Jump)

	mongoTokens := StackTokens.Mongo

	mongoDevTokens := mongoTokens.Development
	mongoProdTokens := mongoTokens.Production

	mongoDevNaming := modules.NewNaming(stack, cfg, mongoDevTokens)
	mongoProdNaming := modules.NewNaming(stack, cfg, mongoProdTokens)

	mongoDevAddrs := CoreSubnets.Mongo.Development
	mongoProdAddrs := CoreSubnets.Mongo.Production

	mongoDevSubnetInput := resources.NewSubnetInput(stack, mongoDevNaming, nil, mongoDevAddrs)
	mongoProdSubnetInput := resources.NewSubnetInput(stack, mongoProdNaming, nil, mongoProdAddrs)

	subnetInputs := make([]vnet.VirtualNetworkSubnet, 3)
	subnetInputs[CoreSubnetIndices.Jump] = jumpSubnetInput
	subnetInputs[CoreSubnetIndices.Mongo.Development] = mongoDevSubnetInput
	subnetInputs[CoreSubnetIndices.Mongo.Production] = mongoProdSubnetInput

	vnet := resources.NewVNet(stack, cfg, naming, rg, CoreAddressSpace, subnetInputs)

	jumpSubnet := resources.GetSubnet(vnet, CoreSubnetIndices.Jump)
	mongoDevSubnet := resources.GetSubnet(vnet, CoreSubnetIndices.Mongo.Development)
	mongoProdSubnet := resources.GetSubnet(vnet, CoreSubnetIndices.Mongo.Production)

	return DefaultCoreDrum{
		StackName_: StackNames.Core,
		Stack_:     stack,
		JumpBeat_: DefaultJumpCoreBeat{
			Naming_: jumpNaming,
			Subnet_: jumpSubnet,
			ASG_:    jumpASG,
			NSG_:    jumpNSG,
		},
		MongoBeats_: DefaultMongoCoreBeats{
			Development_: DefaultMongoCoreBeat{
				Naming_: mongoDevNaming,
				Subnet_: mongoDevSubnet,
			},
			Production_: DefaultMongoCoreBeat{
				Naming_: mongoProdNaming,
				Subnet_: mongoProdSubnet,
			},
		},
	}
}
