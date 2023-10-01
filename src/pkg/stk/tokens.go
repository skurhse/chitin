package stk

import cfg "github.com/transprogrammer/xenia/pkg/config"

type TokensIndex struct {
	Core  string
	Jump  string
	Mongo string
	Dev   string
	Prod  string
}

var Tokens = TokensIndex{
	Core:  "core",
	Jump:  "jump",
	Mongo: "mongo",
	Dev:   "dev",
	Prod:  "prod",
}

func EnrichTokens(cfg cfg.Config, tokens []string) {
	return append(cfg.Name, tokens)
}
