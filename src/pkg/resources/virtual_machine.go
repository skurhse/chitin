package resources

import (
	"fmt"

	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	nic "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/networkinterface"
	rg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/resourcegroup"
	vm "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualmachine"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/cfg"
)

func NewVirtualMachine(stack cdktf.TerraformStack, cfg cfg.Config, naming naming.Naming, rg rg.ResourceGroup, nic nic.NetworkInterface) vm.VirtualMachine {

	storageImageReference := vm.VirtualMachineStorageImageReference{
		Publisher: jsii.String("Canonical"),
		Offer:     jsii.String("0001-com-ubuntu-server-jammy"),
		Sku:       jsii.String("22_04-lts-gen2"),
		Version:   jsii.String("latest"),
	}

	storageOSDisk := vm.VirtualMachineStorageOsDisk{
		Name:            jsii.String("osdisk"),
		CreateOption:    jsii.String("FromImage"),
		ManagedDiskType: jsii.String("Standard_LRS"),
	}

	networkInterfaceIds := []*string{nic.Id()}

	adminId := jsii.String("admin_username")
	adminConfig := cdktf.TerraformVariableConfig{
		Type:        jsii.String("string"),
		Description: jsii.String("administrator username"),
		Sensitive:   jsii.Bool(true),
	}
	adminUsername := cdktf.NewTerraformVariable(stack, adminId, &adminConfig)

	keyId := jsii.String("key_data")
	keyConfig := cdktf.TerraformVariableConfig{
		Type:        jsii.String("string"),
		Description: jsii.String("administrator username"),
		Sensitive:   jsii.Bool(true),
	}
	keyData := cdktf.NewTerraformVariable(stack, keyId, &keyConfig)

	osProfile := vm.VirtualMachineOsProfile{
		ComputerName:  naming.VirtualMachineOutput(),
		AdminUsername: adminUsername.StringValue(),
		AdminPassword: keyData.StringValue(),
	}

	sshPath := fmt.Sprintf("/home/%s/.ssh/authorized_keys", adminUsername.StringValue())

	sshKeys := vm.VirtualMachineOsProfileLinuxConfigSshKeys{
		Path:    &sshPath,
		KeyData: keyData.StringValue(),
	}

	osProfileLinuxConfig := vm.VirtualMachineOsProfileLinuxConfig{
		DisablePasswordAuthentication: jsii.Bool(true),
		SshKeys:                       sshKeys,
	}

	input := vm.VirtualMachineConfig{
		Name:                  naming.VirtualMachineOutput(),
		Location:              cfg.Regions().Primary(),
		ResourceGroupName:     rg.Name(),
		VmSize:                jsii.String("Standard_B2ms"),
		StorageImageReference: &storageImageReference,
		StorageDataDisk:       &storageOSDisk,
		NetworkInterfaceIds:   &networkInterfaceIds,
		OsProfile:             &osProfile,
		OsProfileLinuxConfig:  &osProfileLinuxConfig,
	}

	return vm.NewVirtualMachine(stack, Ids.VirtualMachine, &input)
}
