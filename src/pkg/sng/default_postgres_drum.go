package sng

import "github.com/hashicorp/terraform-cdk-go/cdktf"

type DefaultPostgresMelody struct {
	StackName_ *string
	Stack_     cdktf.TerraformStack
}

func (self DefaultPostgresMelody) StackName() *string {
	return self.StackName_
}

func (self DefaultPostgresMelody) Stack() cdktf.TerraformStack {
	return self.Stack_
}
