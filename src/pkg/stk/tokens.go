package stk

import (
	"github.com/skurhse/xen/pkg/cfg"
)

type TokensIndex struct {
	Core     string
	Jump     string
	Postgres string
	Cluster  string
}

type TokenSetsIndex struct {
	Core     []string
	Jump     []string
	Postgres []string
	Cluster  []string
}

var Tokens = TokensIndex{
	Core:     "core",
	Jump:     "jump",
	Postgres: "postgres",
}

func NewTokenSets(cfg cfg.Config) TokenSetsIndex {
	name := *cfg.Name()

	return TokenSetsIndex{
		Core:     []string{name, Tokens.Core},
		Jump:     []string{name, Tokens.Jump},
		Postgres: []string{name, Tokens.Postgres},
		Cluster:  []string{name, Tokens.Cluster},
	}
}
