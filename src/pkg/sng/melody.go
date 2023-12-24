package sng

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	cnf "github.com/skurhse/chitin/generated/hashicorp/azurerm/dataazurermclientconfig"
)

type Melody interface {
	Stack() cdktf.TerraformStack
	StackName() *string
	Client() cnf.DataAzurermClientConfig
}
