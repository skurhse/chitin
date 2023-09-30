package resources

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	asg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/applicationsecuritygroup"
	rg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/resourcegroup"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/apps"
)

var appSecurityGroupResourceId = jsii.String("application_security_group")

func NewAppSecurityGroup(stack cdktf.TerraformStack, config apps.Config, *naming naming.Naming, rg *resourceGroup rg.ResourceGroup) *asg.ApplicationSecurityGroup {

	id := ResourceIds

	input := &asg.ApplicationSecurityGroupConfig{
		Name:              naming.ApplicationSecurityGroupOutput(),
		Location:          config.Regions().Primary(),
		ResourceGroupName: resourceGroup.Name(),
	}

	return asg.NewApplicationSecurityGroup(stack, id, input)
}
