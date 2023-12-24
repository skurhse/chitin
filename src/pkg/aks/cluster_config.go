package aks

import "github.com/skurhse/chitin/pkg/cfg"

type ClusterConfig interface {
	cfg.Config
	WhitelistIPs() *[]*string
}
