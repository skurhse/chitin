package stacks

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	asg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/applicationsecuritygroup"
	nsg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/networksecuritygroup"
	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/apps"
	"github.com/transprogrammer/xenia/pkg/providers"
	"github.com/transprogrammer/xenia/pkg/resources"
)

type JumpCoreBeat interface {
	CoreBeat
	ASG() *asg.ApplicationSecurityGroup
	NSG() *nsg.NetworkSecurityGroup
}

type DefaultJumpCoreBeat struct {
	Naming_ *naming.Naming
	Subnet_ *vnet.VirtualNetworkSubnetOutputReference
	ASG_    *asg.ApplicationSecurityGroup
	NSG_    *nsg.NetworkSecurityGroup
	VNet_   *vnet.VirtualNetwork
}

func (c DefaultJumpCoreBeat) Naming() *naming.Naming {
	return c.Naming_
}

func (c DefaultJumpCoreBeat) Subnet() *vnet.VirtualNetworkSubnetOutputReference {
	return c.Subnet_
}

func (c DefaultJumpCoreBeat) ASG() *asg.ApplicationSecurityGroup {
	return c.ASG_
}

func (c DefaultJumpCoreBeat) NSG() *nsg.NetworkSecurityGroup {
	return c.NSG_
}

func (c DefaultJumpCoreBeat) VNet() *vnet.VirtualNetwork {
	return c.VNet_
}

type JumpDrum interface {
	StackDrum
}

type DefaultJumpDrum struct {
	StackName_ *string
	Stack_     *cdktf.TerraformStack
}

func (self DefaultJumpDrum) StackName() *string {
	return self.StackName_
}

func (self DefaultJumpDrum) Stack() *cdktf.TerraformStack {
	return self.Stack_
}

func NewJump(app constructs.Construct, cfg apps.Config, core JumpCoreBeat) DefaultJumpDrum {
	stackName := StackNames.Jump
	stack := cdktf.NewTerraformStack(app, stackName)
	providers.NewAzureRM(stack, cfg)

	naming := core.Naming()
	asg := core.ASG()
	nsg := core.NSG()
	subnet := core.Subnet()

	rg := resources.NewResourceGroup(stack, cfg, naming)

	ip := resources.NewPublicIP(stack, cfg, naming, rg)

	nic := resources.NewNIC(stack, cfg, naming, rg, subnet, ip)

	resources.NewNICAssocASG(stack, cfg, nic, asg)
	resources.NewNICAssocNSG(stack, cfg, nic, nsg)

	resources.NewVirtualMachine(stack, cfg, naming, rg, nic)

	return DefaultJumpDrum{
		StackName_: jumpStackName,
		Stack_:     stack,
	}
}
