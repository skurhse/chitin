package res

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	pg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/postgresqlflexibleserver"
	ad "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/postgresqlflexibleserveractivedirectoryadministrator"
	rg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/resourcegroup"
	cc "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/dataazurermclientconfig"
)

func NewPostgresADAdmin(stack cdktf.TerraformStack, rg rg.ResourceGroup, client cc.DataAzurermClientConfig, server pg.PostgresqlFlexibleServer) ad.PostgresqlFlexibleServerActiveDirectoryAdministrator {

	input := ad.PostgresqlFlexibleServerActiveDirectoryAdministratorConfig{
		ServerName:        server.Name(),
		ResourceGroupName: rg.Name(),
		TenantId:          client.TenantId(),
	}

	return ad.NewPostgresqlFlexibleServerActiveDirectoryAdministrator(stack, Ids.PostgresAdmin, &input)
}
