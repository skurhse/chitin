package linuxvirtualmachine

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	lvm "github.com/skurhse/chitin/generated/hashicorp/azurerm/linuxvirtualmachine"
	nic "github.com/skurhse/chitin/generated/hashicorp/azurerm/networkinterface"
	rg "github.com/skurhse/chitin/generated/hashicorp/azurerm/resourcegroup"
	"github.com/skurhse/chitin/generated/naming"
	"github.com/skurhse/chitin/pkg/cfg"
)

func NewVirtualMachine(stack cdktf.TerraformStack, cfg cfg.Config, naming naming.Naming, rg rg.ResourceGroup, nic nic.NetworkInterface) lvm.LinuxVirtualMachine {

	imageReference := lvm.LinuxVirtualMachineSourceImageReference{
		Publisher: jsii.String("Canonical"),
		Offer:     jsii.String("0001-com-ubuntu-server-jammy"),
		Sku:       jsii.String("22_04-lts-gen2"),
		Version:   jsii.String("latest"),
	}

	osDisk := lvm.LinuxVirtualMachineOsDisk{
		Name:               jsii.String("osdisk"),
		Caching:            jsii.String("ReadWrite"),
		StorageAccountType: jsii.String("Standard_LRS"),
	}

	networkInterfaceIds := []*string{nic.Id()}

	adminId := jsii.String("admin_username")
	adminConfig := cdktf.TerraformVariableConfig{
		Type:        jsii.String("string"),
		Description: jsii.String("administrator username"),
		Sensitive:   jsii.Bool(true),
	}
	adminUsername := cdktf.NewTerraformVariable(stack, adminId, &adminConfig)

	keyId := jsii.String("public_key")
	keyConfig := cdktf.TerraformVariableConfig{
		Type:        jsii.String("string"),
		Description: jsii.String("public key"),
		Sensitive:   jsii.Bool(true),
	}
	keyData := cdktf.NewTerraformVariable(stack, keyId, &keyConfig)

	sshKey := lvm.LinuxVirtualMachineAdminSshKey{
		Username:  adminUsername.StringValue(),
		PublicKey: keyData.StringValue(),
	}
	sshKeys := []*lvm.LinuxVirtualMachineAdminSshKey{&sshKey}

	input := lvm.LinuxVirtualMachineConfig{
		AdminUsername:        adminUsername.StringValue(),
		Name:                 naming.LinuxVirtualMachineOutput(),
		Location:             cfg.Regions().Primary(),
		ResourceGroupName:    rg.Name(),
		Size:                 jsii.String("Standard_B2ms"),
		SourceImageReference: &imageReference,
		OsDisk:               &osDisk,
		NetworkInterfaceIds:  &networkInterfaceIds,
		AdminSshKey:          &sshKeys,
	}

	return lvm.NewLinuxVirtualMachine(stack, Ids.VirtualMachine, &input)
}
