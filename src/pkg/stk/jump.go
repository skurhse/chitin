package stk

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/transprogrammer/xenia/pkg/mod"
	"github.com/transprogrammer/xenia/pkg/prv"
	"github.com/transprogrammer/xenia/pkg/res"
)

func NewJump(app constructs.Construct, cfg JumpConfig, core JumpCoreBeat, tokens []string) DefaultJumpDrum {
	name := NewStackName(tokens)

	stk := cdktf.NewTerraformStack(app, name)
	prv.NewAzureRM(stk)

	naming := core.Naming()
	asg := core.ASG()
	nsg := core.NSG()
	subnet := core.Subnet()

	rg := res.NewResourceGroup(stk, cfg, naming)

	jumpName := mod.NewNaming(stk, tokenSets.Jump)
	jumpASG := res.NewASG(stk, cfg, jumpName, rg)
	jumpSecurityRule := res.NewSSHSecurityRule(cfg.WhitelistIPs(), jumpASG)
	jumpNSG := res.NewNSG(stk, cfg, jumpName, rg, jumpSecurityRule)
	jumpSubnet := res.NewSubnet(stk, jumpName, rg, vnet, CoreSubnetAddrs.Jump, Tokens.Jump)
	res.NewSubnetNSGAssoc(stk, jumpSubnet, jumpNSG, Tokens.Jump)

	ip := res.NewPublicIP(stk, cfg, naming, rg)

	nic := res.NewNIC(stk, cfg, naming, rg, subnet, ip)

	res.NewNICAssocASG(stk, cfg, nic, asg)
	res.NewNICAssocNSG(stk, cfg, nic, nsg)

	res.NewVirtualMachine(stk, cfg, naming, rg, nic)

	// res.NewAdminGroup(stk, cfg, naming, rg)

	return DefaultJumpDrum{
		StackName_: name,
		Stack_:     stk,
	}
}
