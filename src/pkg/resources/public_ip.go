package resources

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	ip "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/public_ip"
	"github.com/transprogrammer/xenia/generated/hashicorp/azurerm/resourcegroup"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/apps"
)

func NewPublicIP(stack cdktf.TerraformStack, config config.Config, naming *naming.Naming, rg *resourcegroup.ResourceGroup) *ip.PublicIp {
	id := ResourceIds.PublicIP

	input := &ip.PublicIpConfig{
		Name:                 (*naming).PublicIpOutput(),
		Location:             config.Regions().Primary(),
		ResourceGroupName:    (*rg).Name(),
		Sku:                  jsii.String("Basic"),
		AllocationMethod:     jsii.String("Dynamic"),
		IpVersion:            jsii.String("IPv4"),
		DomainNameLabel:      apps.AppName,
		IdleTimeoutInMinutes: jsii.Number(4),
	}

	ip := ip.NewPublicIp(stack, id, input)

	return &ip
}
