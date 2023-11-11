package stk

import "github.com/transprogrammer/xenia/pkg/cfg"

type ClusterConfig interface {
	cfg.Config
	WhitelistIPs() *[]*string
}
