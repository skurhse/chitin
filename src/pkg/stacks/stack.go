package stacks

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewStack(scope constructs.Construct, name *string) cdktf.TerraformStack {
	return cdktf.NewTerraformStack(scope, name)
}
