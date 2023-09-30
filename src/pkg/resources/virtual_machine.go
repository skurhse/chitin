package resources

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	nic "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/networkinterface"
	rg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/resourcegroup"
	vm "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualmachine"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/apps"
)

func NewVirtualMachine(stack cdktf.TerraformStack, config config.Config, naming *naming.Naming, rg *rg.ResourceGroup, nic *nic.NetworkInterface) *vm.VirtualMachine {
	id := ResourceIds.VirtualMachine

	storageImageReference := &vm.VirtualMachineStorageImageReference{
		Publisher: config.VirtualMachine.Image.Publisher,
		Offer:     config.VirtualMachine.Image.Offer,
		Sku:       config.VirtualMachine.Image.Sku,
		Version:   config.VirtualMachine.Image.Version,
	}

	storageOSDisk := &vm.VirtualMachineStorageOsDisk{
		Name:            jsii.String("osdisk"),
		CreateOption:    jsii.String("FromImage"),
		ManagedDiskType: config.VirtualMachine.StorageAccountType,
	}

	networkInterfaceIds := &[]*string{nic.Id()}

	osProfile := &vm.VirtualMachineOsProfile{
		ComputerName:  naming.VirtualMachineOutput(),
		AdminUsername: config.VirtualMachine.AdminUsername,
		AdminPassword: config.VirtualMachine.SSHPublicKey,
	}

	sshKeys := &vm.VirtualMachineOsProfileLinuxConfigSshKeys{
		Path:    jsii.String("/home/" + *config.VirtualMachine.AdminUsername + "/.ssh/authorized_keys"),
		KeyData: config.VirtualMachine.SSHPublicKey,
	}

	osProfileLinuxConfig := &vm.VirtualMachineOsProfileLinuxConfig{
		DisablePasswordAuthentication: jsii.Bool(true),
		SshKeys:                       sshKeys,
	}

	input := &vm.VirtualMachineConfig{
		Name:                  naming.VirtualMachineOutput(),
		Location:              config.Regions.Primary,
		ResourceGroupName:     rg.Name(),
		VmSize:                config.VirtualMachine.Size,
		StorageImageReference: storageImageReference,
		StorageDataDisk:       storageOSDisk,
		NetworkInterfaceIds:   networkInterfaceIds,
		OsProfile:             osProfile,
		OsProfileLinuxConfig:  osProfileLinuxConfig,
	}

	return vm.NewVirtualMachine(stack, id, input)
}
