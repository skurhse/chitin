package stk

import "github.com/skurhse/xen/pkg/cfg"

type JumpConfig interface {
	cfg.Config
	WhitelistIPs() *[]*string
}
