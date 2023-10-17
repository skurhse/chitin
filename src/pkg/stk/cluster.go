package stk

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/cfg"
	"github.com/transprogrammer/xenia/pkg/providers"
	"github.com/transprogrammer/xenia/pkg/resources"
)

type ClusterDrum interface {
	Drum
}

type DefaultClusterDrum struct {
	StackName_ *string
	Stack_     cdktf.TerraformStack
}

func (self DefaultClusterDrum) StackName() *string {
	return self.StackName_
}

func (self DefaultClusterDrum) Stack() cdktf.TerraformStack {
	return self.Stack_
}

type ClusterConfig interface {
	cfg.Config
	WhitelistIPs() *[]*string
}

type ClusterCoreBeat interface {
	CoreBeat
}

type DefaultClusterCoreBeat struct {
	Naming_ naming.Naming
	Subnet_ vnet.VirtualNetworkSubnetOutputReference
}

func (c DefaultClusterCoreBeat) Naming() naming.Naming {
	return c.Naming_
}

func (c DefaultClusterCoreBeat) Subnet() vnet.VirtualNetworkSubnetOutputReference {
	return c.Subnet_
}

func NewCluster(app constructs.Construct, cfg ClusterConfig, core ClusterCoreBeat, jump ClusterJumpBeat, tokens []string) DefaultClusterDrum {
	subnet := core.Subnet()

	name := NewName(tokens)

	stk := cdktf.NewTerraformStack(app, name)
	providers.NewAzureRM(stk)

	naming := core.Naming()

	resources.NewResourceGroup(stk, cfg, naming)
	resources.NewCluster(stk, cfg, naming, rg, subnet)

	return DefaultClusterDrum{
		StackName_: name,
		Stack_:     stk,
	}
}
