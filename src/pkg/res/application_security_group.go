package resources

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	asg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/applicationsecuritygroup"
	rg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/resourcegroup"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/cfg"
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
