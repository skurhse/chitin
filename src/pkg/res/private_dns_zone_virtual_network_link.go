package res

import (
	"fmt"

	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	pdnsz "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/privatednszone"
	pdnszvnl "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/privatednszonevirtualnetworklink"
	rg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/resourcegroup"
	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/cfg"
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
