package dataazurermclientconfig

import (
	"fmt"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	cnf "github.com/skurhse/chitin/generated/hashicorp/azurerm/dataazurermclientconfig"
	"github.com/skurhse/chitin/pkg/res"
)

func NewDataAzurermClientConfig(stack cdktf.TerraformStack) cnf.DataAzurermClientConfig {

	input := cnf.DataAzurermClientConfigConfig{}

	name := fmt.Sprintf("%s_1", res.Ids.ClientConfig)

	return cnf.NewDataAzurermClientConfig(stack, &name, &input)
}
