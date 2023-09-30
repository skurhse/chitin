package stacks

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/apps"
	cfg "github.com/transprogrammer/xenia/pkg/config"
	"github.com/transprogrammer/xenia/pkg/providers"
	"github.com/transprogrammer/xenia/pkg/resources"

	dbacct "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/cosmosdbaccount"
	db "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/cosmosdbmongodatabase"
	pdnsz "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/privatednszone"
	pdnszvnl "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/privatednszonevirtualnetworklink"
	pe "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/privateendpoint"
	"github.com/transprogrammer/xenia/generated/hashicorp/azurerm/resourcegroup"
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
	Development_ DefaultMongoCoreBeat
	Production_  DefaultMongoCoreBeat
}

type DefaultMongoCoreBeat struct {
	Naming_ *naming.Naming
	Subnet_ *vnet.VirtualNetworkSubnetOutputReference
}

func (c DefaultMongoCoreBeats) Development() MongoCoreBeat {
	return c.Development()
}

func (c DefaultMongoCoreBeats) Production() MongoCoreBeat {
	return c.Production()
}

func (c DefaultMongoCoreBeat) Naming() *naming.Naming {
	return c.Naming_
}

func (c DefaultMongoCoreBeat) Subnet() *vnet.VirtualNetworkSubnetOutputReference {
	return c.Subnet_
}

func NewMongo(app constructs.Construct, cfg cfg.Config, core MongoCoreBeat) DefaultMongoDrum {
	stackName := StackNames.Mongo

	stk := cdktf.NewTerraformStack(app, stackName)
	providers.NewAzureRM(stk, cfg)

	naming := core.Naming()
	subnet := core.Subnet()

	rg := resource.NewResourceGroup(stk, cfg, naming)

	acct := NewMongoAccount(stk, cfg, env, naming, rg)

	NewMongoDatabase(stk, cfg, env, naming, acct)
	NewPrivateEndpoint(stk, cfg, env, naming, rg, subnet)

	return stack
}

func NewMongoAccount(stack cdktf.TerraformStack, cfg cfg.AppConfig, env cfg.MongoEnvironment, naming naming.Naming, rg resourcegroup.ResourceGroup) {

	input := dbacct.CosmosdbAccountConfig{
		Name:                       naming.CosmosdbAccountOutput(),
		Location:                   cfg.Regions().Primary(),
		ResourceGroupName:          rg.Name(),
		Kind:                       jsii.String("MongoDB"),
		OfferType:                  jsii.String("Standard"),
		MongoServerVersion:         env.ServerVersion(),
		PublicNetworkAccessEnabled: jsii.Bool(false),
		ConsistencyPolicy: &dbacct.CosmosdbAccountConsistencyPolicy{
			ConsistencyLevel: env.ConsistencyLevel(),
		},
		GeoLocation: &[]*dbacct.CosmosdbAccountGeoLocation{
			{
				Location:         cfg.Regions().Secondary(),
				FailoverPriority: jsii.Number(0),
				ZoneRedundant:    jsii.Bool(false),
			},
		},
		Capabilities: &[]*dbacct.CosmosdbAccountCapabilities{
			{
				Name: jsii.String("DisabledRateLimitingResponses"),
			},
			{
				Name: jsii.String("EnableServerless"),
			},
		},
	}

	return dbacct.NewCosmosdbAccount(stack, resources.Ids.NewCosmosDBAccount, &input)
}

func NewMongoDatabase(stk cdktf.TerraformStack, cfg cfg.Config, naming naming.Naming, acct dbacct.CosmosdbAccount) db.MongoDatabase {

	id := resources.Ids().MongoDatabase

	input := db.NewMongoDatabaseConfig{
		AccountName:       acct.Name(),
		Name:              naming.MongoDatabaseOutputs(),
		resourceGroupName: rg.Name(),
	}

	return db.NewMongoDatabase(stk, id, &input)
}

func NewMongoPrivateEndpoint(stk cdktf.TerraformStack, cfg cfg.AppConfiguration, naming naming.Naming, rg rg.ResourceGroup, acct dbacct.CosmosdbAccount, subnet vnet.VirtualNetworkSubnetOutputReference) pe.PrivateEndpoint {

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

	return pe.NewPrivateEndpoint(stk, id, &input)
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
