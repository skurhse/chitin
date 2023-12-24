package cre

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/skurhse/chitin/pkg/mod"
	"github.com/skurhse/chitin/pkg/prv"
	"github.com/skurhse/chitin/pkg/sng"
	rg "github.skurhse/chitin/pkg/res/resource_group"
	vnet "github.skurhse/chitin/pkg/res/virtual_network"
)

const (
	coreAddr = "10.0.0.0/16"
	jumpAddr = "10.1.0.0./24"
	pgAddr   = "10.2.0.0./24"
)

var CoreAddrSpace = []*string{jsii.String(coreAddr)}

var CoreSubnetAddrs = CoreSubnetAddrsIndex{
	Jump:     jsii.String(jumpAddr),
	Postgres: jsii.String(pgAddr),
}

func NewCore(scope constructs.Construct, cfg CoreConfig, token string) DefaultCoreMelody {

	name := sng.NewName(cfg.Name(), token)

	stk := sng.NewStack(scope, name)
	prv.NewAzureRM(stk)

	naming := mod.NewNaming(stk, cfg.Name(), token)

	rg := rg.NewResourceGroup(stk, cfg, naming)
	clnt := cnf.NewDataAzurermClientConfig(stk)
	vnet := vnet.NewVirtualNetwork(stk, naming, rg, CoreAddrSpace, token)

	return DefaultCoreMelody{
		Stack_:          stk,
		StackName_:      name,
		Client_:         clnt,
		VirtualNetwork_: vnet,
	}
}
