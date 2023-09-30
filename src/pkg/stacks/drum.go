package stacks

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StackDrum interface {
	StackName() *string
	Stack() *cdktf.TerraformStack
}
