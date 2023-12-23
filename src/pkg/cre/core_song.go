package cre

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/skurhse/xen/pkg/mod"
	"github.com/skurhse/xen/pkg/prv"
	"github.com/skurhse/xen/pkg/res"
	"github.com/skurhse/xen/pkg/sng"
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

	rg := res.NewResourceGroup(stk, cfg, naming)
	clnt := res.NewDataAzurermClientConfig(stk)
	vnet := res.NewVirtualNetwork(stk, naming, rg, CoreAddrSpace, token)

	return DefaultCoreMelody{
		Stack_:          stk,
		StackName_:      name,
		Client_:         clnt,
		VirtualNetwork_: vnet,
	}
}
