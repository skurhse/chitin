package sng

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/skurhse/xen/pkg/mod"
	"github.com/skurhse/xen/pkg/prv"
	"github.com/skurhse/xen/pkg/res"
)

func NewJump(app constructs.Construct, cfg JumpConfig, core JumpCoreTune, tokens []string) DefaultJumpMelody {
	name := NewStackName(tokens)

	stk := cdktf.NewTerraformStack(app, name)
	prv.NewAzureRM(stk)

	naming := core.Naming()
	asg := core.ASG()
	nsg := core.NSG()
	vnet := core.VirtualNetwork()

	rg := res.NewResourceGroup(stk, cfg, naming)

	jumpName := mod.NewNaming(stk, tokens)
	jumpASG := res.NewASG(stk, cfg, jumpName, rg)
	jumpSecurityRule := res.NewSSHSecurityRule(cfg.WhitelistIPs(), jumpASG)
	jumpNSG := res.NewNSG(stk, cfg, jumpName, rg, jumpSecurityRule)
	jumpSubnet := res.NewSubnet(stk, jumpName, rg, vnet, CoreSubnetAddrs.Jump, Tokens.Jump)
	res.NewSubnetNSGAssoc(stk, jumpSubnet, jumpNSG, Tokens.Jump)

	ip := res.NewPublicIP(stk, cfg, naming, rg)

	nic := res.NewNIC(stk, cfg, naming, rg, jumpSubnet, ip)

	res.NewNICAssocASG(stk, cfg, nic, asg)
	res.NewNICAssocNSG(stk, cfg, nic, nsg)

	res.NewVirtualMachine(stk, cfg, naming, rg, nic)

	// res.NewAdminGroup(stk, cfg, naming, rg)

	return DefaultJumpMelody{
		StackName_: name,
		Stack_:     stk,
	}
}
