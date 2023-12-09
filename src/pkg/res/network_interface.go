package res

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	asg "github.com/skurhse/xen/generated/hashicorp/azurerm/applicationsecuritygroup"
	nic "github.com/skurhse/xen/generated/hashicorp/azurerm/networkinterface"
	nicasg "github.com/skurhse/xen/generated/hashicorp/azurerm/networkinterfaceapplicationsecuritygroupassociation"
	nicnsg "github.com/skurhse/xen/generated/hashicorp/azurerm/networkinterfacesecuritygroupassociation"
	nsg "github.com/skurhse/xen/generated/hashicorp/azurerm/networksecuritygroup"
	"github.com/skurhse/xen/generated/hashicorp/azurerm/publicip"
	"github.com/skurhse/xen/generated/hashicorp/azurerm/resourcegroup"
	sn "github.com/skurhse/xen/generated/hashicorp/azurerm/subnet"
	"github.com/skurhse/xen/generated/naming"
	"github.com/skurhse/xen/pkg/cfg"
)

func NewNIC(stack cdktf.TerraformStack, cfg cfg.Config, naming naming.Naming, rg resourcegroup.ResourceGroup, subnet sn.Subnet, ip publicip.PublicIp) nic.NetworkInterface {
	ipConfig := nic.NetworkInterfaceIpConfiguration{
		Name:                       jsii.String("ipcfg"),
		Primary:                    jsii.Bool(true),
		SubnetId:                   subnet.Id(),
		PublicIpAddressId:          ip.Id(),
		PrivateIpAddressAllocation: jsii.String("Dynamic"),
	}

	input := &nic.NetworkInterfaceConfig{
		Name:              naming.NetworkInterfaceOutput(),
		Location:          cfg.Regions().Primary(),
		ResourceGroupName: rg.Name(),
		IpConfiguration:   &[]*nic.NetworkInterfaceIpConfiguration{&ipConfig},
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
