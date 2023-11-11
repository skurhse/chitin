package stk

type JumpCoreBeat interface {
	CoreBeat
	ASG() asg.ApplicationSecurityGroup
	NSG() nsg.NetworkSecurityGroup
}
