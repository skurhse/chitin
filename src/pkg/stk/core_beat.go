package stk

import (
	sn "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/subnet"
	"github.com/transprogrammer/xenia/generated/naming"
)

type CoreBeat interface {
	Naming() naming.Naming
	Subnet() sn.Subnet
}
