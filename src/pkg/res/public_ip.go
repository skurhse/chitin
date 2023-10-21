package resources

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	ip "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/publicip"
	"github.com/transprogrammer/xenia/generated/hashicorp/azurerm/resourcegroup"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/cfg"
)

func NewPublicIP(stack cdktf.TerraformStack, cfg cfg.Config, naming naming.Naming, rg resourcegroup.ResourceGroup) ip.PublicIp {
	input := ip.PublicIpConfig{
		Name:                 naming.PublicIpOutput(),
		Location:             cfg.Regions().Primary(),
		ResourceGroupName:    rg.Name(),
		Sku:                  jsii.String("Basic"),
		AllocationMethod:     jsii.String("Dynamic"),
		IpVersion:            jsii.String("IPv4"),
		DomainNameLabel:      cfg.Name(),
		IdleTimeoutInMinutes: jsii.Number(4),
	}

	return ip.NewPublicIp(stack, Ids.PublicIP, &input)
}
