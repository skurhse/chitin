package res

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	pg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/postgresqlflexibleserver"
	dns "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/privatednszone"
	nl "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/privatednszonevirtualnetworklink"
	rg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/resourcegroup"
	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/cfg"
)

// PORT: Using postgres server naming as stand-in. <>
// ???: Add flexi-server resource definition to naming fork? <rbt>

func NewPostgresFlexibleServer(stack cdktf.TerraformStack, cfg cfg.Config, naming naming.Naming, rg rg.ResourceGroup, subnet vnet.VirtualNetworkSubnetOutputReference, zone dns.PrivateDnsZone, vnetLink nl.PrivateDnsZoneVirtualNetworkLink, tenantId *string) pg.PostgresqlFlexibleServer {

	serverVersion := jsii.String("15")
	storageMB := jsii.Number(32768)
	skuName := jsii.String("B_Standard_B1ms")

	auth := &pg.PostgresqlFlexibleServerAuthentication{
		ActiveDirectoryAuthEnabled: jsii.Bool(true),
		PasswordAuthEnabled:        jsii.Bool(false),
		TenantId:                   tenantId,
	}

	input := pg.PostgresqlFlexibleServerConfig{
		Name:              naming.PostgresqlServerOutput(),
		ResourceGroupName: rg.Name(),
		Location:          cfg.Regions().Primary(),
		Version:           serverVersion,
		DelegatedSubnetId: subnet.Id(),
		PrivateDnsZoneId:  zone.Id(),
		Authentication:    auth,
		StorageMb:         storageMB,
		SkuName:           skuName,

		DependsOn: &[]cdktf.ITerraformDependable{vnetLink},
	}

	return pg.NewPostgresqlFlexibleServer(stack, Ids.PostgresFlexibleServer, &input)
}
