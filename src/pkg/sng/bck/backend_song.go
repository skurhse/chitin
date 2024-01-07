package bck

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/skurhse/chitin/pkg/mod"
	"github.com/skurhse/chitin/pkg/prv"
	clnt "github.com/skurhse/chitin/pkg/res/dataazurermclientconfig"
	rg "github.com/skurhse/chitin/pkg/res/resourcegroup"
	"github.com/skurhse/chitin/pkg/sng"
)

func NewBackend(scope constructs.Construct, cfg BackendConfig, token string) BackendMelody {
	name := sng.NewStackName(cfg.Name(), token)

	stk := sng.NewStack(scope, name)
	prv.NewAzureRM(stk)

	naming := mod.NewNaming(stk, cfg.Name(), token)

	rg := rg.NewResourceGroup(stk, cfg, naming)
	clnt := clnt.NewDataAzurermClientConfig(stk)
	stor := stor.NewStorageAccount(stk, cfg, naming)
	vnet := vnet.NewVirtualNetwork(stk, naming, rg, BackendAddrSpace, token)

	return BackendMelody{
		Stack_:          stk,
		StackName_:      name,
		Client_:         clnt,
		VirtualNetwork_: vnet,
	}
}
