package stk

import "github.com/hashicorp/terraform-cdk-go/cdktf"

type DefaultClusterDrum struct {
	StackName_ *string
	Stack_     cdktf.TerraformStack
}

func (self DefaultClusterDrum) StackName() *string {
	return self.StackName_
}

func (self DefaultClusterDrum) Stack() cdktf.TerraformStack {
	return self.Stack_
}
