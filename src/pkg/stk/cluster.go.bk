package stk

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/cfg"
	"github.com/transprogrammer/xenia/pkg/prv"
	"github.com/transprogrammer/xenia/pkg/res"
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

	adminGroup := jump.AdminGroup()

	name := NewName(tokens)

	stk := cdktf.NewTerraformStack(app, name)
	prv.NewAzureRM(stk)

	naming := core.Naming()

	rg := res.NewResourceGroup(stk, cfg, naming)
	res.NewCluster(stk, cfg, naming, rg, subnet, adminGroup)

	return DefaultClusterDrum{
		StackName_: name,
		Stack_:     stk,
	}
}
