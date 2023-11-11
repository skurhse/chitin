package main

import "github.com/hashicorp/terraform-cdk-go/cdktf"

type DefaultPostgresDrum struct {
	StackName_ *string
	Stack_     cdktf.TerraformStack
}

func (self DefaultPostgresDrum) StackName() *string {
	return self.StackName_
}

func (self DefaultPostgresDrum) Stack() cdktf.TerraformStack {
	return self.Stack_
}
