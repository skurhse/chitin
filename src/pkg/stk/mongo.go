package stk

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/apps"
	"github.com/transprogrammer/xenia/pkg/cfg"
	"github.com/transprogrammer/xenia/pkg/providers"
	"github.com/transprogrammer/xenia/pkg/resources"

	dbacct "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/cosmosdbaccount"
	db "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/cosmosdbmongodatabase"
	pdnsz "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/privatednszone"
	pdnszvnl "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/privatednszonevirtualnetworklink"
	pe "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/privateendpoint"
	rg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/resourcegroup"
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
	Development() MongoCoreBeat
	Production() MongoCoreBeat
}

type MongoCoreBeat interface {
	CoreBeat
}

type DefaultMongoCoreBeats struct {
	Dev_  DefaultMongoCoreBeat
	Prod_ DefaultMongoCoreBeat
}

type DefaultMongoCoreBeat struct {
	Naming_ naming.Naming
	Subnet_ vnet.VirtualNetworkSubnetOutputReference
}

func (c DefaultMongoCoreBeats) Development() MongoCoreBeat {
	return c.Development()
}

func (c DefaultMongoCoreBeats) Production() MongoCoreBeat {
	return c.Production()
}

func (c DefaultMongoCoreBeat) Naming() naming.Naming {
	return c.Naming_
}

func (c DefaultMongoCoreBeat) Subnet() vnet.VirtualNetworkSubnetOutputReference {
	return c.Subnet_
}

func NewMongo(scope constructs.Construct, cfg cfg.Config, core MongoCoreBeat, tokens []string) DefaultMongoDrum {
	name := NewName(tokens)

	stack := cdktf.NewTerraformStack(scope, name)
	providers.NewAzureRM(stack)

	naming := core.Naming()
	subnet := core.Subnet()

	rg := resources.NewResourceGroup(stack, cfg, naming)

	acct := NewMongoAccount(stack, cfg, naming, rg)

	NewMongoDatabase(stack, cfg, naming, acct)
	NewPrivateEndpoint(stack, cfg, naming, rg, subnet)

	return DefaultMongoDrum{
		StackName_: name,
		Stack_:     stack,
	}
}

func NewMongoDatabase(stack cdktf.TerraformStack, cfg cfg.Config, naming naming.Naming, acct dbacct.CosmosdbAccount) db.MongoDatabase {

	id := resources.Ids().MongoDatabase

	input := db.NewMongoDatabaseConfig{
		AccountName:       acct.Name(),
		Name:              naming.MongoDatabaseOutputs(),
		resourceGroupName: rg.Name(),
	}

	return db.NewMongoDatabase(stack, id, &input)
}

func NewMongoPrivateEndpoint(stack cdktf.TerraformStack, cfg cfg.AppConfiguration, naming naming.Naming, rg rg.ResourceGroup, acct dbacct.CosmosdbAccount, subnet vnet.VirtualNetworkSubnetOutputReference) pe.PrivateEndpoint {

	id := resources.Ids().PrivateEndpoint()

	conn := pe.PrivateEndpointPrivateServiceConnection{
		Name:                        jsii.String("cosmosdb"),
		PrivateConnectionResourceId: acct.Id(),
		SubresourceNames:            &[]*string{jsii.String("MongoDB")},
	}

	input := pe.PrivateEndpointConfig{
		Name:                     naming.PrivateEndpointOutput(),
		Location:                 apps.Regions().Primary(),
		ResourceGroupName:        rg.Name(),
		SubnetId:                 subnet.Id(),
		PrivateServiceConnection: &conn,
	}

	return pe.NewPrivateEndpoint(stack, id, &input)
}

var PrivateDNSZone pdnsz.PrivateDnsZone = pdnsz.NewPrivateDnsZone(Stk, Ids.PrivateDNSZone, &pdnsz.PrivateDnsZoneConfig{
	Name:              jsii.String("privatelink.mongo.cosmos.azure.com"),
	ResourceGroupName: Rg.Name(),
})

var PrivateDNSZoneVirtualNetworkLink pdnszvnl.PrivateDnsZoneVirtualNetworkLink = pdnszvnl.NewPrivateDnsZoneVirtualNetworkLink(Stk, Ids.PrivateDNSZoneVirtualNetworkLink, &pdnszvnl.PrivateDnsZoneVirtualNetworkLinkConfig{
	Name:                jsii.String(fmt.Sprintf("%s-vnetlink", *MongoNaming.PrivateDnsZoneOutput())),
	ResourceGroupName:   Rg.Name(),
	PrivateDnsZoneName:  PrivateDNSZone.Name(),
	VirtualNetworkId:    VNet.Id(),
	RegistrationEnabled: jsii.Bool(true),
})
