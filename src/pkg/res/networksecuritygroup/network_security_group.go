package networksecuritygroup

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	asg "github.com/skurhse/chitin/generated/hashicorp/azurerm/applicationsecuritygroup"
	nsg "github.com/skurhse/chitin/generated/hashicorp/azurerm/networksecuritygroup"
	rg "github.com/skurhse/chitin/generated/hashicorp/azurerm/resourcegroup"
	"github.com/skurhse/chitin/generated/naming"
	"github.com/skurhse/chitin/pkg/cfg"
)

func NewNSG(stack cdktf.TerraformStack, cfg cfg.Config, naming naming.Naming, rg rg.ResourceGroup, securityRule nsg.NetworkSecurityGroupSecurityRule) nsg.NetworkSecurityGroup {

	id := res.Ids.NetworkSecurityGroup

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
