package stacks

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/modules"
	"github.com/transprogrammer/xenia/pkg/providers"
	"github.com/transprogrammer/xenia/pkg/resources"
)

type CoreDrum interface {
	StackDrum
	JumpBeat() JumpCoreBeat
	MongoBeats() MongoCoreBeats
}

type DefaultCoreDrum struct {
	StackName_  *string
	Stack_      *cdktf.TerraformStack
	JumpBeat_   DefaultJumpCoreBeat
	MongoBeats_ DefaultMongoCoreBeats
}

type CoreBeat interface {
	Naming() *naming.Naming
	Subnet() *vnet.VirtualNetworkSubnetOutputReference
}

func (c DefaultCoreDrum) StackName() *string {
	return c.StackName_
}

func (c DefaultCoreDrum) Stack() *cdktf.TerraformStack {
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

func NewCore(app constructs.Construct) DefaultCoreDrum {
	stackName := StackNames.Core
	stackTokens := StackTokens.Core

	stack := cdktf.NewTerraformStack(app, stackName)
	providers.NewAzureRM(stack, config)

	naming := modules.NewNaming(stack, config, stackTokens)

	rg := resources.NewResourceGroup(stack, config, naming)

	jumpTokens := StackTokens.Jump
	jumpNaming := modules.NewNaming(stack, config, jumpTokens)

	jumpASG := resources.NewAppSecurityGroup(stack, config, jumpNaming, rg)

	jumpSecurityRule := resources.NewSSHSecurityRule(config, jumpASG)

	jumpNSG := resources.NewNSG(stack, config, jumpNaming, rg, jumpSecurityRule)

	jumpSubnetInput := resources.NewSubnetInput(stack, jumpNaming, jumpNSG, CoreSubnets.Jump)

	mongoTokens := StackTokens.Mongo

	mongoDevTokens := mongoTokens.Development
	mongoProdTokens := mongoTokens.Production

	mongoDevNaming := modules.NewNaming(stack, config, mongoDevTokens)
	mongoProdNaming := modules.NewNaming(stack, config, mongoProdTokens)

	mongoDevAddrs := CoreSubnets.Mongo.Development
	mongoProdAddrs := CoreSubnets.Mongo.Production

	mongoDevSubnetInput := resources.NewSubnetInput(stack, mongoDevNaming, nil, mongoDevAddrs)
	mongoProdSubnetInput := resources.NewSubnetInput(stack, mongoProdNaming, nil, mongoProdAddrs)

	subnetInputs := make([]vnet.VirtualNetworkSubnet, 3)
	subnetInputs[CoreSubnetIndices.Jump] = jumpSubnetInput
	subnetInputs[CoreSubnetIndices.Mongo.Development] = mongoDevSubnetInput
	subnetInputs[CoreSubnetIndices.Mongo.Production] = mongoProdSubnetInput

	vnet := resources.NewVNet(stack, config, naming, rg, CoreAddressSpace, subnetInputs)

	jumpSubnet := resources.GetSubnet(vnet, CoreSubnetIndices.Jump)
	mongoDevSubnet := resources.GetSubnet(vnet, CoreSubnetIndices.Mongo.Development)
	mongoProdSubnet := resources.GetSubnet(vnet, CoreSubnetIndices.Mongo.Production)

	return DefaultCoreDrum{
		StackName_: stackName,
		Stack_:     &stack,
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
