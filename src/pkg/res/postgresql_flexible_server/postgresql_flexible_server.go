package res

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	cnf "github.com/skurhse/xen/generated/hashicorp/azurerm/dataazurermclientconfig"
	pg "github.com/skurhse/xen/generated/hashicorp/azurerm/postgresqlflexibleserver"
	dns "github.com/skurhse/xen/generated/hashicorp/azurerm/privatednszone"
	nl "github.com/skurhse/xen/generated/hashicorp/azurerm/privatednszonevirtualnetworklink"
	rg "github.com/skurhse/xen/generated/hashicorp/azurerm/resourcegroup"
	sn "github.com/skurhse/xen/generated/hashicorp/azurerm/subnet"
	"github.com/skurhse/xen/generated/naming"
	"github.com/skurhse/xen/pkg/cfg"
)

// PORT: Using postgres server naming as stand-in. <>
// ???: Add flexi-server resource definition to naming fork? <rbt>

func NewPostgresFlexibleServer(stack cdktf.TerraformStack, cfg cfg.Config, naming naming.Naming, rg rg.ResourceGroup, subnet sn.Subnet, zone dns.PrivateDnsZone, vnetLink nl.PrivateDnsZoneVirtualNetworkLink, client cnf.DataAzurermClientConfig) pg.PostgresqlFlexibleServer {

	serverVersion := jsii.String("15")
	storageMB := jsii.Number(32768)
	skuName := jsii.String("B_Standard_B1ms")
	backupEnabled := jsii.Bool(true)

	auth := &pg.PostgresqlFlexibleServerAuthentication{
		ActiveDirectoryAuthEnabled: jsii.Bool(true),
		PasswordAuthEnabled:        jsii.Bool(false),
		TenantId:                   client.TenantId(),
	}

	input := pg.PostgresqlFlexibleServerConfig{
		Name:                      naming.PostgresqlServerOutput(),
		ResourceGroupName:         rg.Name(),
		Location:                  cfg.Regions().Primary(),
		Version:                   serverVersion,
		DelegatedSubnetId:         subnet.Id(),
		PrivateDnsZoneId:          zone.Id(),
		Authentication:            auth,
		StorageMb:                 storageMB,
		SkuName:                   skuName,
		GeoRedundantBackupEnabled: backupEnabled,

		DependsOn: &[]cdktf.ITerraformDependable{vnetLink},
	}

	return pg.NewPostgresqlFlexibleServer(stack, Ids.PostgresFlexibleServer, &input)
}
