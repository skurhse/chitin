package sng

import "github.com/hashicorp/terraform-cdk-go/cdktf"

type DefaultClusterMelody struct {
	StackName_ *string
	Stack_     cdktf.TerraformStack
}

func (self DefaultClusterMelody) StackName() *string {
	return self.StackName_
}

func (self DefaultClusterMelody) Stack() cdktf.TerraformStack {
	return self.Stack_
}
