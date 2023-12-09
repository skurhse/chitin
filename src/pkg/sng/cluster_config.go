package sng

import "github.com/skurhse/xen/pkg/cfg"

type ClusterConfig interface {
	cfg.Config
	WhitelistIPs() *[]*string
}
