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

type JumpDrum interface {
	Drum
}

type DefaultJumpDrum struct {
	StackName_ *string
	Stack_     cdktf.TerraformStack
}

func (self DefaultJumpDrum) StackName() *string {
	return self.StackName_
}

func (self DefaultJumpDrum) Stack() cdktf.TerraformStack {
	return self.Stack_
}

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

func NewJump(app constructs.Construct, cfg stacks.Config, core JumpCoreBeat) DefaultJumpDrum {
	stkName := StackNames.Jump
	stk := cdktf.NewTerraformStack(app, stkName)
	providers.NewAzureRM(stk, cfg)

	naming := core.Naming()
	asg := core.ASG()
	nsg := core.NSG()
	subnet := core.Subnet()

	rg := resources.NewResourceGroup(stk, cfg, naming)

	ip := resources.NewPublicIP(stk, cfg, naming, rg)

	nic := resources.NewNIC(stk, cfg, naming, rg, subnet, ip)

	resources.NewNICAssocASG(stk, cfg, nic, asg)
	resources.NewNICAssocNSG(stk, cfg, nic, nsg)

	resources.NewVirtualMachine(stk, cfg, naming, rg, nic)

	return DefaultJumpDrum{
		StackName_: stkName,
		Stack_:     stk,
	}
}
