package stk

import "github.com/transprogrammer/xenia/pkg/cfg"

const (
	CoreToken    = "core"
	JumpToken    = "jump"
	MongoToken   = "mongo"
	DevToken     = "dev"
	ProdToken    = "prod"
	ClusterToken = "cluster"
)

type Tokens struct {
	Core    []string
	Jump    []string
	Mongo   MongoTokens
	Cluster []string
}

type MongoTokens struct {
	Dev  []string
	Prod []string
}

func NewTokens(cfg cfg.Config) Tokens {
	name := *cfg.Name()

	return Tokens{
		Core: []string{name, CoreToken},
		Jump: []string{name, JumpToken},
		Mongo: MongoTokens{
			Dev:  []string{name, MongoToken, DevToken},
			Prod: []string{name, MongoToken, ProdToken},
		},
		Cluster: []string{name, ClusterToken},
	}
}
