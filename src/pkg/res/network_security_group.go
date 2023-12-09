package res

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	asg "github.com/skurhse/xen/generated/hashicorp/azurerm/applicationsecuritygroup"
	nsg "github.com/skurhse/xen/generated/hashicorp/azurerm/networksecuritygroup"
	rg "github.com/skurhse/xen/generated/hashicorp/azurerm/resourcegroup"
	"github.com/skurhse/xen/generated/naming"
	"github.com/skurhse/xen/pkg/cfg"
)

func NewNSG(stack cdktf.TerraformStack, cfg cfg.Config, naming naming.Naming, rg rg.ResourceGroup, securityRule nsg.NetworkSecurityGroupSecurityRule) nsg.NetworkSecurityGroup {

	id := Ids.NetworkSecurityGroup

	input := nsg.NetworkSecurityGroupConfig{
		Name:              naming.NetworkSecurityGroupOutput(),
		Location:          cfg.Regions().Primary(),
		ResourceGroupName: rg.Name(),
		SecurityRule:      &[]*nsg.NetworkSecurityGroupSecurityRule{&securityRule},
	}

	return nsg.NewNetworkSecurityGroup(stack, id, &input)
}

func NewSSHSecurityRule(ips *[]*string, asg asg.ApplicationSecurityGroup) nsg.NetworkSecurityGroupSecurityRule {

	groupIds := &[]*string{asg.Id()}

	return nsg.NetworkSecurityGroupSecurityRule{
		Name:                                   jsii.String("SSH"),
		Description:                            jsii.String("Allow SSH"),
		Priority:                               jsii.Number(100),
		Direction:                              jsii.String("Inbound"),
		Access:                                 jsii.String("Allow"),
		Protocol:                               jsii.String("Tcp"),
		DestinationPortRange:                   jsii.String("22"),
		DestinationApplicationSecurityGroupIds: groupIds,
		SourceAddressPrefixes:                  ips,
		SourcePortRange:                        jsii.String("*"),
	}
}
