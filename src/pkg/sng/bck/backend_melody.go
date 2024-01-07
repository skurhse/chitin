package bck

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	cnf "github.com/skurhse/chitin/generated/hashicorp/azurerm/dataazurermclientconfig"
	vnet "github.com/skurhse/chitin/generated/hashicorp/azurerm/virtualnetwork"
)

type BackendMelody struct {
	Stack_          cdktf.TerraformStack
	StackName_      *string
	Client_         cnf.DataAzurermClientConfig
	VirtualNetwork_ vnet.VirtualNetwork
}

func (c BackendMelody) Stack() cdktf.TerraformStack {
	return c.Stack_
}

func (c BackendMelody) StackName() *string {
	return c.StackName_
}

func (c BackendMelody) Client() cnf.DataAzurermClientConfig {
	return c.Client_
}

func (c BackendMelody) VirtualNetwork() vnet.VirtualNetwork {
	return c.VirtualNetwork_
}
