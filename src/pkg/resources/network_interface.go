package resources

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	asg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/applicationsecuritygroup"
	nic "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/networkinterface"
	nicasg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/networkinterfaceapplicationsecuritygroupassociation"
	nicnsg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/networkinterfacesecuritygroupassociation"
	nsg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/networksecuritygroup"
	"github.com/transprogrammer/xenia/generated/hashicorp/azurerm/publicip"
	"github.com/transprogrammer/xenia/generated/hashicorp/azurerm/resourcegroup"
	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/apps"
)

func NewNIC(stack cdktf.TerraformStack, config stacks.Config, naming *naming.Naming, rg *resourcegroup.ResourceGroup, subnet *vnet.VirtualNetworkSubnetOutputReference, ip *publicip.PublicIp) *nic.NetworkInterface {
	id := ResourceIds.NetworkInterface

	ipConfig := nic.NetworkInterfaceIpConfiguration{
		Name:                    jsii.String("ipconfig"),
		Primary:                 jsii.Bool(true),
		SubnetId:                (*subnet).Id(),
		PublicIpAddressId:       (*ip).Id(),
		PrivateIpAddressVersion: jsii.String("Dynamic"),
	}

	input := &nic.NetworkInterfaceConfig{
		Name:              (*naming).NetworkInterfaceOutput(),
		Location:          config.Regions().Primary(),
		ResourceGroupName: (*rg).Name(),
		IpConfiguration:   ipConfig,
	}

	nic := nic.NewNetworkInterface(stack, id, input)

	return &nic
}

func NewNICAssocASG(stack cdktf.TerraformStack, config stacks.Config, nic *nic.NetworkInterface, asg *asg.ApplicationSecurityGroup) *nicasg.NetworkInterfaceApplicationSecurityGroupAssociation {
	id := ResourceIds.NetworkInterfaceApplicationSecurityGroupAssociation

	input := &nicasg.NetworkInterfaceApplicationSecurityGroupAssociationConfig{
		NetworkInterfaceId:         (*nic).Id(),
		ApplicationSecurityGroupId: (*asg).Id(),
	}

	assoc := nicasg.NewNetworkInterfaceApplicationSecurityGroupAssociation(stack, id, input)

	return &assoc
}

func NewNICAssocNSG(stack cdktf.TerraformStack, config stacks.Config, nic *nic.NetworkInterface, nsg *nsg.NetworkSecurityGroup) *nicnsg.NetworkInterfaceSecurityGroupAssociation {
	id := ResourceIds.NetworkInterfaceNetworkSecurityGroupAssociation

	input := &nicnsg.NetworkInterfaceSecurityGroupAssociationConfig{
		NetworkInterfaceId:     (*nic).Id(),
		NetworkSecurityGroupId: (*nsg).Id(),
	}

	assoc := nicnsg.NewNetworkInterfaceSecurityGroupAssociation(stack, id, input)

	return &assoc
}
