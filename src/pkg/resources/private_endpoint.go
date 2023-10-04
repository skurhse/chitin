package resources

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	dbacct "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/cosmosdbaccount"
	pdnsz "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/privatednszone"
	pe "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/privateendpoint"
	rg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/resourcegroup"
	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/cfg"
)

// TODO: Modularize. <rbt>

func NewPrivateEndpoint(stack cdktf.TerraformStack, cfg cfg.Config, naming naming.Naming, rg rg.ResourceGroup, acct dbacct.CosmosdbAccount, subnet vnet.VirtualNetworkSubnetOutputReference, zone pdnsz.PrivateDnsZone) pe.PrivateEndpoint {

	privateServiceConnection := pe.PrivateEndpointPrivateServiceConnection{
		Name:                        jsii.String("cosmosdb"),
		PrivateConnectionResourceId: acct.Id(),
		SubresourceNames:            &[]*string{jsii.String("MongoDB")},
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

	return pe.NewPrivateEndpoint(stack, Ids.PrivateEndpoint, &input)
}
