package sng

import "github.com/hashicorp/terraform-cdk-go/cdktf"

type DefaultJumpMelody struct {
	StackName_ *string
	Stack_     cdktf.TerraformStack
}

func (self DefaultJumpMelody) StackName() *string {
	return self.StackName_
}

func (self DefaultJumpMelody) Stack() cdktf.TerraformStack {
	return self.Stack_
}
