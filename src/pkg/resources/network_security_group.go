package resources

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	asg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/applicationsecuritygroup"
	nsg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/networksecuritygroup"
	rg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/resourcegroup"
	"github.com/transprogrammer/xenia/generated/naming"
	"honnef.co/go/tools/config"
)

func NewNSG(stack cdktf.TerraformStack, config config.Config, naming naming.Naming, rg rg.ResourceGroup, securityRule nsg.NetworkSecurityGroupSecurityRule) nsg.NetworkSecurityGroup {

	id := Ids.NetworkSecurityGroup

	input := nsg.NetworkSecurityGroupConfig{
		Name:              naming.NetworkSecurityGroupOutput(),
		Location:          config.Regions().Primary(),
		ResourceGroupName: rg.Name(),
		SecurityRule:      &securityRule,
	}

	return nsg.NewNetworkSecurityGroup(stack, id, &input)
}

func NewSSHSecurityRule(cfg config.Config, asg *asg.ApplicationSecurityGroup) nsg.NetworkSecurityGroupSecurityRule {

	groupIds := &[]*string{(*asg).Id()}

	srcPrefixes := cfg.SSHSourceAddressPrefixes()

	return nsg.NetworkSecurityGroupSecurityRule{
		Name:                                   jsii.String("SSH"),
		Description:                            jsii.String("Allow SSH"),
		Priority:                               jsii.Number(100),
		Direction:                              jsii.String("Inbound"),
		Access:                                 jsii.String("Allow"),
		Protocol:                               jsii.String("Tcp"),
		DestinationPortRange:                   jsii.String("22"),
		DestinationApplicationSecurityGroupIds: groupIds,
		SourceAddressPrefixes:                  srcPrefixes,
		SourcePortRange:                        jsii.String("*"),
	}
}
