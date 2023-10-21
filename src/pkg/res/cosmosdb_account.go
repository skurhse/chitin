package res

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	dbacct "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/cosmosdbaccount"
	"github.com/transprogrammer/xenia/generated/hashicorp/azurerm/resourcegroup"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/cfg"
)

func NewCosmosDBMongoAccount(stack cdktf.TerraformStack, cfg cfg.Config, naming naming.Naming, rg resourcegroup.ResourceGroup) dbacct.CosmosdbAccount {

	consistencyPolicy := dbacct.CosmosdbAccountConsistencyPolicy{
		ConsistencyLevel: jsii.String("Eventual"),
	}

	geoLocation := []*dbacct.CosmosdbAccountGeoLocation{
		{
			Location:         cfg.Regions().Secondary(),
			FailoverPriority: jsii.Number(0),
			ZoneRedundant:    jsii.Bool(false),
		},
	}

	acctCapabilities := []*dbacct.CosmosdbAccountCapabilities{
		{
			Name: jsii.String("DisabledRateLimitingResponses"),
		},
		{
			Name: jsii.String("EnableServerless"),
		},
	}

	input := dbacct.CosmosdbAccountConfig{
		Name:                       naming.CosmosdbAccountOutput(),
		Location:                   cfg.Regions().Primary(),
		ResourceGroupName:          rg.Name(),
		Kind:                       jsii.String("MongoDB"),
		OfferType:                  jsii.String("Standard"),
		MongoServerVersion:         jsii.String("4.2"),
		PublicNetworkAccessEnabled: jsii.Bool(false),
		ConsistencyPolicy:          &consistencyPolicy,
		GeoLocation:                &geoLocation,
		Capabilities:               &acctCapabilities,
	}

	return dbacct.NewCosmosdbAccount(stack, Ids.CosmosDBAccount, &input)
}
