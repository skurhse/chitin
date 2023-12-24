package publicip

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	ip "github.com/skurhse/xen/generated/hashicorp/azurerm/publicip"
	"github.com/skurhse/xen/generated/hashicorp/azurerm/resourcegroup"
	"github.com/skurhse/xen/generated/naming"
	"github.com/skurhse/xen/pkg/cfg"
)

func NewPublicIP(stack cdktf.TerraformStack, cfg cfg.Config, naming naming.Naming, rg resourcegroup.ResourceGroup) ip.PublicIp {

	input := ip.PublicIpConfig{
		Name:                 naming.PublicIpOutput(),
		Location:             cfg.Regions().Primary(),
		ResourceGroupName:    rg.Name(),
		Sku:                  jsii.String("Basic"),
		AllocationMethod:     jsii.String("Dynamic"),
		IpVersion:            jsii.String("IPv4"),
		DomainNameLabel:      jsii.String(cfg.Name()),
		IdleTimeoutInMinutes: jsii.Number(4),
	}

	return ip.NewPublicIp(stack, Ids.PublicIP, &input)
}
