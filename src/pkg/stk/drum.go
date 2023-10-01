package stk

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Drum interface {
	StackName() *string
	Stack() cdktf.TerraformStack
}
