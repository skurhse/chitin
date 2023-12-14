package cre

import "github.com/hashicorp/terraform-cdk-go/cdktf"

type DefaultCoreMelody struct {
	Name_  *string
	Stack_ cdktf.TerraformStack
}

func (c DefaultCoreMelody) StackName() *string {
	return c.Name_
}

func (c DefaultCoreMelody) Stack() cdktf.TerraformStack {
	return c.Stack_
}
