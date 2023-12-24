package privatednszonevirtualnetworklink

import (
	"fmt"

	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	pdnsz "github.com/skurhse/chitin/generated/hashicorp/azurerm/privatednszone"
	pdnszvnl "github.com/skurhse/chitin/generated/hashicorp/azurerm/privatednszonevirtualnetworklink"
	rg "github.com/skurhse/chitin/generated/hashicorp/azurerm/resourcegroup"
	vnet "github.com/skurhse/chitin/generated/hashicorp/azurerm/virtualnetwork"
	"github.com/skurhse/chitin/generated/naming"
	"github.com/skurhse/chitin/pkg/cfg"
)

func NewPrivateDNSZoneVNetLink(stack cdktf.TerraformStack, cfg cfg.Config, naming naming.Naming, rg rg.ResourceGroup, zone pdnsz.PrivateDnsZone, vnet vnet.VirtualNetwork) pdnszvnl.PrivateDnsZoneVirtualNetworkLink {

	name := fmt.Sprintf("%s-vnetlink", *naming.PrivateDnsZoneOutput())

	input := pdnszvnl.PrivateDnsZoneVirtualNetworkLinkConfig{
		Name:                &name,
		ResourceGroupName:   rg.Name(),
		PrivateDnsZoneName:  zone.Name(),
		VirtualNetworkId:    vnet.Id(),
		RegistrationEnabled: jsii.Bool(true),
	}

	return pdnszvnl.NewPrivateDnsZoneVirtualNetworkLink(stack, Ids.PrivateDNSZoneVNetLink, &input)
}
