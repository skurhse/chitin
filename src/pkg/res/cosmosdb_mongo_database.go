package resources

import (
	"fmt"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	dbacct "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/cosmosdbaccount"
	db "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/cosmosdbmongodatabase"
	rg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/resourcegroup"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/cfg"
)

func NewCosmosDBMongoDatabase(stack cdktf.TerraformStack, cfg cfg.Config, naming naming.Naming, rg rg.ResourceGroup, acct dbacct.CosmosdbAccount) db.CosmosdbMongoDatabase {
	// TODO: Update naming module <>
	name := fmt.Sprintf("%s-db", *naming.CosmosdbAccountOutput())

	input := db.CosmosdbMongoDatabaseConfig{
		AccountName:       acct.Name(),
		Name:              &name,
		ResourceGroupName: rg.Name(),
	}

	return db.NewCosmosdbMongoDatabase(stack, Ids.CosmosDBMongoDatabase, &input)
}
