package stk

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/cfg"
	"github.com/transprogrammer/xenia/pkg/providers"
	"github.com/transprogrammer/xenia/pkg/res"

	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
)

type MongoDrum interface {
	Drum
}

type DefaultMongoDrum struct {
	StackName_ *string
	Stack_     cdktf.TerraformStack
}

func (self DefaultMongoDrum) StackName() *string {
	return self.StackName_
}

func (self DefaultMongoDrum) Stack() cdktf.TerraformStack {
	return self.Stack_
}

type MongoConfig interface {
	cfg.Config
}

type MongoCoreBeats interface {
	Dev() MongoCoreBeat
	Prod() MongoCoreBeat
}

type MongoCoreBeat interface {
	CoreBeat
	VNet() vnet.VirtualNetwork
}

type DefaultMongoCoreBeats struct {
	Dev_  DefaultMongoCoreBeat
	Prod_ DefaultMongoCoreBeat
}

type DefaultMongoCoreBeat struct {
	Naming_ naming.Naming
	Subnet_ vnet.VirtualNetworkSubnetOutputReference
	VNet_   vnet.VirtualNetwork
}

func (c DefaultMongoCoreBeats) Dev() MongoCoreBeat {
	return c.Dev_
}

func (c DefaultMongoCoreBeats) Prod() MongoCoreBeat {
	return c.Prod_
}

func (c DefaultMongoCoreBeat) Naming() naming.Naming {
	return c.Naming_
}

func (c DefaultMongoCoreBeat) Subnet() vnet.VirtualNetworkSubnetOutputReference {
	return c.Subnet_
}

func (c DefaultMongoCoreBeat) VNet() vnet.VirtualNetwork {
	return c.VNet_
}

func NewMongo(scope constructs.Construct, cfg cfg.Config, core MongoCoreBeat, tokens []string) DefaultMongoDrum {
	name := NewName(tokens)

	stack := cdktf.NewTerraformStack(scope, name)
	providers.NewAzureRM(stack)

	naming := core.Naming()
	subnet := core.Subnet()
	vnet := core.VNet()

	rg := res.NewResourceGroup(stack, cfg, naming)

	acct := res.NewCosmosDBMongoAccount(stack, cfg, naming, rg)

	res.NewCosmosDBMongoDatabase(stack, cfg, naming, rg, acct)

	zone := res.NewPrivateDNSZone(stack, rg)
	res.NewPrivateDNSZoneVNetLink(stack, cfg, naming, rg, zone, vnet)

	res.NewPrivateEndpoint(stack, cfg, naming, rg, acct, subnet, zone)

	return DefaultMongoDrum{
		StackName_: name,
		Stack_:     stack,
	}
}
