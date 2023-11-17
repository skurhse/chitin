package stk

import "github.com/transprogrammer/xenia/pkg/cfg"

const (
	CoreToken     = "core"
	JumpToken     = "jump"
	PostgresToken = "postgres"
	ClusterToken  = "cluster"
)

type Tokens struct {
	Core     []string
	Jump     []string
	Postgres []string
	Cluster  []string
}

func NewTokens(cfg cfg.Config) Tokens {
	name := *cfg.Name()

	return Tokens{
		Core:     []string{name, CoreToken},
		Jump:     []string{name, JumpToken},
		Postgres: []string{name, PostgresToken},
		Cluster:  []string{name, ClusterToken},
	}
}
