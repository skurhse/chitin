package sng

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Melody interface {
	Name() *string
	Stack() cdktf.TerraformStack
}
