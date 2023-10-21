package stk

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	asg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/applicationsecuritygroup"
	nsg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/networksecuritygroup"
	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/cfg"
	"github.com/transprogrammer/xenia/pkg/prv"
	"github.com/transprogrammer/xenia/pkg/res"
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

type JumpConfig interface {
	cfg.Config
	WhitelistIPs() *[]*string
}

type JumpCoreBeat interface {
	CoreBeat
	ASG() asg.ApplicationSecurityGroup
	NSG() nsg.NetworkSecurityGroup
}

type DefaultJumpCoreBeat struct {
	Naming_ naming.Naming
	Subnet_ vnet.VirtualNetworkSubnetOutputReference
	ASG_    asg.ApplicationSecurityGroup
	NSG_    nsg.NetworkSecurityGroup
	VNet_   vnet.VirtualNetwork
}

func (c DefaultJumpCoreBeat) Naming() naming.Naming {
	return c.Naming_
}

func (c DefaultJumpCoreBeat) Subnet() vnet.VirtualNetworkSubnetOutputReference {
	return c.Subnet_
}

func (c DefaultJumpCoreBeat) ASG() asg.ApplicationSecurityGroup {
	return c.ASG_
}

func (c DefaultJumpCoreBeat) NSG() nsg.NetworkSecurityGroup {
	return c.NSG_
}

func (c DefaultJumpCoreBeat) VNet() vnet.VirtualNetwork {
	return c.VNet_
}

type JumpBeat interface{}

func NewJump(app constructs.Construct, cfg JumpConfig, core JumpCoreBeat, tokens []string) DefaultJumpDrum {
	name := NewName(tokens)

	stk := cdktf.NewTerraformStack(app, name)
	prv.NewAzureRM(stk)

	naming := core.Naming()
	asg := core.ASG()
	nsg := core.NSG()
	subnet := core.Subnet()

	rg := res.NewResourceGroup(stk, cfg, naming)

	ip := res.NewPublicIP(stk, cfg, naming, rg)

	nic := res.NewNIC(stk, cfg, naming, rg, subnet, ip)

	res.NewNICAssocASG(stk, cfg, nic, asg)
	res.NewNICAssocNSG(stk, cfg, nic, nsg)

	res.NewVirtualMachine(stk, cfg, naming, rg, nic)

	res.New

	return DefaultJumpDrum{
		StackName_: name,
		Stack_:     stk,
	}
}
