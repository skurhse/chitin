package main

import (
	"fmt"
	"log"

	"github.com/transprogrammer/xenia/pkg/apps"
	"github.com/transprogrammer/xenia/pkg/aspects"
	"github.com/transprogrammer/xenia/pkg/cfg"
	"github.com/transprogrammer/xenia/pkg/stacks"
	"github.com/transprogrammer/xenia/pkg/stk"
)

func main() {
	app := apps.App
	tokens := stk.Tokens

	cfg, err := cfg.Load()
	if err != nil {
		err = fmt.Errorf("load config: %w", err)
		log.Fatal(err)
	}

	core := stacks.NewCore(app, cfg, tokens.Core)
	jumpBeat := core.JumpBeat()
	mongoBeats := core.MongoBeats()
	mongoDevBeat := mongoBeats.Development()
	mongoProdBeat := mongoBeats.Production()

	drums := [4]stacks.Drum{
		core,
		stacks.NewJump(app, cfg, jumpBeat, tokens.Jump),
		stacks.NewMongo(app, cfg, mongoDevBeat, tokens.Mongo, tokens.Dev),
		stacks.NewMongo(app, cfg, mongoProdBeat, tokens.Mongo, tokens.Prod),
	}

	for _, drum := range drums {
		aspects.AddTags(drum, cfg)
	}

	app.Synth()
}
