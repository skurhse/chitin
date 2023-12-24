package sng

import "github.com/skurhse/chitin/pkg/cfg"

type JumpConfig interface {
	cfg.Config
	WhitelistIPs() *[]*string
}
