package dataazurermclientconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	cnf "github.com/skurhse/chitin/generated/hashicorp/azurerm/dataazurermclientconfig"
)

func NewDataAzurermClientConfig(stack cdktf.TerraformStack) cnf.DataAzurermClientConfig {

	input := cnf.DataAzurermClientConfigConfig{}

	return cnf.NewDataAzurermClientConfig(stack, Ids.ClientConfig, &input)
}
