package res


pg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/postgresqlflexibleserveractivedirectoryadministrator"

func NewPostgresADAdmin(stack cdktf.TerraformStack, cfg cfg.Config, naming naming.Naming, rg rg.ResourceGroup) pg.PostgresqlFlexibleServerActiveDirectoryAdministrator {


	input := pg.PostgresqlFlexibleServerActiveDirectoryAdministratorConfig{
		ServerName:
		ResourceGroupName: rg.Name(),
		TenantId: 
	}

}
