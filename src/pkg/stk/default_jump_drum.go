package stk

import "github.com/hashicorp/terraform-cdk-go/cdktf"

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
