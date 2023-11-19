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

	adminGroup := res.NewAdminGroup(stk, cfg, naming, rg)

	return DefaultJumpDrum{
		StackName_: name,
		Stack_:     stk,
	}
}
