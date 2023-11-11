package stk

import "github.com/transprogrammer/xenia/pkg/cfg"

type JumpConfig interface {
	cfg.Config
	WhitelistIPs() *[]*string
}
