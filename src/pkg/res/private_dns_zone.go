package res

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	pdnsz "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/privatednszone"
	rg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/resourcegroup"
)

func NewPrivateDNSZone(stack cdktf.TerraformStack, rg rg.ResourceGroup) pdnsz.PrivateDnsZone {

	input := pdnsz.PrivateDnsZoneConfig{
		Name:              jsii.String("privatelink.postgres.cosmos.azure.com"),
		ResourceGroupName: rg.Name(),
	}

	return pdnsz.NewPrivateDnsZone(stack, Ids.PrivateDNSZone, &input)
}
