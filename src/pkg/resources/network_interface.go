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
	"github.com/transprogrammer/xenia/pkg/cfg"
)

func NewNIC(stack cdktf.TerraformStack, cfg cfg.Config, naming naming.Naming, rg resourcegroup.ResourceGroup, subnet vnet.VirtualNetworkSubnetOutputReference, ip publicip.PublicIp) nic.NetworkInterface {
	ipConfig := nic.NetworkInterfaceIpConfiguration{
		Name:                    jsii.String("ipcfg"),
		Primary:                 jsii.Bool(true),
		SubnetId:                subnet.Id(),
		PublicIpAddressId:       ip.Id(),
		PrivateIpAddressVersion: jsii.String("Dynamic"),
	}

	input := &nic.NetworkInterfaceConfig{
		Name:              naming.NetworkInterfaceOutput(),
		Location:          cfg.Regions().Primary(),
		ResourceGroupName: rg.Name(),
		IpConfiguration:   ipConfig,
	}

	return nic.NewNetworkInterface(stack, Ids.NetworkInterface, input)
}

func NewNICAssocASG(stack cdktf.TerraformStack, cfg cfg.Config, nic nic.NetworkInterface, asg asg.ApplicationSecurityGroup) nicasg.NetworkInterfaceApplicationSecurityGroupAssociation {
	id := Ids.NetworkInterfaceApplicationSecurityGroupAssociation

	input := &nicasg.NetworkInterfaceApplicationSecurityGroupAssociationConfig{
		NetworkInterfaceId:         nic.Id(),
		ApplicationSecurityGroupId: asg.Id(),
	}

	return nicasg.NewNetworkInterfaceApplicationSecurityGroupAssociation(stack, id, input)
}

func NewNICAssocNSG(stack cdktf.TerraformStack, cfg cfg.Config, nic nic.NetworkInterface, nsg nsg.NetworkSecurityGroup) nicnsg.NetworkInterfaceSecurityGroupAssociation {
	id := Ids.NetworkInterfaceNetworkSecurityGroupAssociation

	input := &nicnsg.NetworkInterfaceSecurityGroupAssociationConfig{
		NetworkInterfaceId:     nic.Id(),
		NetworkSecurityGroupId: nsg.Id(),
	}

	return nicnsg.NewNetworkInterfaceSecurityGroupAssociation(stack, id, input)
}
