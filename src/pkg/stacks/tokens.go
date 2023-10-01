package stacks

type TokenGenerator func(config config.Config, name string) []string 
type EnvTokenGenerator func(config config.Config, env envs.env, name string) []string 

type TokenGeneratorsIndex struct {
	Core TokenGenerator
	Jump TokenGenerator
	Mongo EnvTokenGenerator
}

typ
