package stk

import (
	asg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/applicationsecuritygroup"
	nsg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/networksecuritygroup"
)

type JumpCoreBeat interface {
	CoreBeat
	ASG() asg.ApplicationSecurityGroup
	NSG() nsg.NetworkSecurityGroup
}
