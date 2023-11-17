package main

cnf "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/dataazurermclientconfig"

func newClientConfig(stack cdktf.TerraformStack) cnf.ClientConfig {

	input := cnf.DataAzurermClientConfigConfig{}

	return cnf.NewClientConfig(stack, Id.ClientConfig, &input)
}
