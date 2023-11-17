package stk

import "github.com/hashicorp/terraform-cdk-go/cdktf"

type DefaultCoreDrum struct {
	StackName_    *string
	Stack_        cdktf.TerraformStack
	JumpBeat_     DefaultJumpCoreBeat
	PostgresBeat_ DefaultPostgresCoreBeat
	ClusterBeat_  DefaultClusterCoreBeat
}

func (c DefaultCoreDrum) StackName() *string {
	return c.StackName_
}

func (c DefaultCoreDrum) Stack() cdktf.TerraformStack {
	return c.Stack_
}

func (c DefaultCoreDrum) JumpBeat() JumpCoreBeat {
	return c.JumpBeat_
}

func (c DefaultCoreDrum) PostgresBeat() PostgresCoreBeat {
	return PostgresCoreBeat(c.PostgresBeat_)
}

func (c DefaultCoreDrum) ClusterBeat() ClusterCoreBeat {
	return ClusterCoreBeat(c.ClusterBeat_)
}
