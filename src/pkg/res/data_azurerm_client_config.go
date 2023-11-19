package res

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	cnf "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/dataazurermclientconfig"
)

func NewDataAzurermClientConfig(stack cdktf.TerraformStack) cnf.DataAzurermClientConfig {

	input := cnf.DataAzurermClientConfigConfig{}

	return cnf.NewDataAzurermClientConfig(stack, Ids.ClientConfig, &input)
}