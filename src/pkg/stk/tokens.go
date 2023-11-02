package stk

import "github.com/transprogrammer/xenia/pkg/cfg"

const (
	CoreToken    = "core"
	JumpToken    = "jump"
	PostgresToken   = "postgres"
	DevToken     = "dev"
	ProdToken    = "prod"
	ClusterToken = "cluster"
)

type Tokens struct {
	Core    []string
	Jump    []string
	Postgres   PostgresTokens
	Cluster []string
}

type PostgresTokens struct {
	Dev  []string
	Prod []string
}

func NewTokens(cfg cfg.Config) Tokens {
	name := *cfg.Name()

	return Tokens{
		Core: []string{name, CoreToken},
		Jump: []string{name, JumpToken},
		Postgres: PostgresTokens{
			Dev:  []string{name, PostgresToken, DevToken},
			Prod: []string{name, PostgresToken, ProdToken},
		},
		Cluster: []string{name, ClusterToken},
	}
}
