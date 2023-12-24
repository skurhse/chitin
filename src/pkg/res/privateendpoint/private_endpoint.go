package privateendpoint

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	dbacct "github.com/skurhse/chitin/generated/hashicorp/azurerm/cosmosdbaccount"
	pdnsz "github.com/skurhse/chitin/generated/hashicorp/azurerm/privatednszone"
	pe "github.com/skurhse/chitin/generated/hashicorp/azurerm/privateendpoint"
	rg "github.com/skurhse/chitin/generated/hashicorp/azurerm/resourcegroup"
	sn "github.com/skurhse/chitin/generated/hashicorp/azurerm/subnet"
	"github.com/skurhse/chitin/generated/naming"
	"github.com/skurhse/chitin/pkg/cfg"
)

// TODO: Modularize. <rbt>

func NewPrivateEndpoint(stack cdktf.TerraformStack, cfg cfg.Config, naming naming.Naming, rg rg.ResourceGroup, acct dbacct.CosmosdbAccount, subnet sn.Subnet, zone pdnsz.PrivateDnsZone) pe.PrivateEndpoint {

	privateServiceConnection := pe.PrivateEndpointPrivateServiceConnection{
		Name:                        jsii.String("cosmosdb"),
		PrivateConnectionResourceId: acct.Id(),
		SubresourceNames:            &[]*string{jsii.String("PostgresDB")},
		IsManualConnection:          jsii.Bool(false),
	}

	privateDNSZoneIds := []*string{zone.Id()}

	privateDNSZoneGroup := pe.PrivateEndpointPrivateDnsZoneGroup{
		Name:              naming.PrivateDnsZoneGroupOutput(),
		PrivateDnsZoneIds: &privateDNSZoneIds,
	}

	input := pe.PrivateEndpointConfig{
		Name:                     naming.PrivateEndpointOutput(),
		Location:                 cfg.Regions().Primary(),
		ResourceGroupName:        rg.Name(),
		SubnetId:                 subnet.Id(),
		PrivateServiceConnection: &privateServiceConnection,
		PrivateDnsZoneGroup:      &privateDNSZoneGroup,
	}

	return pe.NewPrivateEndpoint(stack, res.Ids.PrivateEndpoint, &input)
}
