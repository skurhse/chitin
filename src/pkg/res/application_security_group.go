package res

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	asg "github.com/skurhse/xen/generated/hashicorp/azurerm/applicationsecuritygroup"
	rg "github.com/skurhse/xen/generated/hashicorp/azurerm/resourcegroup"
	"github.com/skurhse/xen/generated/naming"
	"github.com/skurhse/xen/pkg/cfg"
)

var appSecurityGroupResourceId = jsii.String("application_security_group")

func NewASG(stack cdktf.TerraformStack, cfg cfg.Config, naming naming.Naming, rg rg.ResourceGroup) asg.ApplicationSecurityGroup {
	input := asg.ApplicationSecurityGroupConfig{
		Name:              naming.ApplicationSecurityGroupOutput(),
		Location:          cfg.Regions().Primary(),
		ResourceGroupName: rg.Name(),
	}

	return asg.NewApplicationSecurityGroup(stack, Ids.AppSecGroup, &input)
}
