package stk

import "github.com/skurhse/xen/pkg/cfg"

type CoreConfig interface {
	cfg.Config
	JumpConfig
	PostgresConfig
}
