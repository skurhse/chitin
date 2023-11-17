package stk

import "github.com/transprogrammer/xenia/pkg/cfg"

type CoreConfig interface {
	cfg.Config
	JumpConfig
	PostgresConfig
}
