package cfg

type TokensIndex struct {
	Backend  string
	Core     string
	Jump     string
	Postgres string
	Cluster  string
}

type TokenSetsIndex struct {
	Backend  string
	Core     []string
	Jump     []string
	Postgres []string
	Cluster  []string
}

var Tokens = TokensIndex{
	Backend:  "backend",
	Core:     "core",
	Jump:     "jump",
	Postgres: "postgres",
}

func NewTokenSets(cfg Config) TokenSetsIndex {
	name := cfg.Name()

	return TokenSetsIndex{
		Core:     []string{name, Tokens.Core},
		Jump:     []string{name, Tokens.Jump},
		Postgres: []string{name, Tokens.Postgres},
		Cluster:  []string{name, Tokens.Cluster},
	}
}
