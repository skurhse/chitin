package cre

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	cnf "github.com/skurhse/chitin/generated/hashicorp/azurerm/dataazurermclientconfig"
	vnet "github.com/skurhse/chitin/generated/hashicorp/azurerm/virtualnetwork"
)

type DefaultCoreMelody struct {
	Stack_          cdktf.TerraformStack
	StackName_      *string
	Client_         cnf.DataAzurermClientConfig
	VirtualNetwork_ vnet.VirtualNetwork
}

func (c DefaultCoreMelody) Stack() cdktf.TerraformStack {
	return c.Stack_
}

func (c DefaultCoreMelody) StackName() *string {
	return c.StackName_
}

func (c DefaultCoreMelody) Client() cnf.DataAzurermClientConfig {
	return c.Client_
}

func (c DefaultCoreMelody) VirtualNetwork() vnet.VirtualNetwork {
	return c.VirtualNetwork_
}
