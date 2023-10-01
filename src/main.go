package main

import (
	"fmt"
	"log"

	"github.com/transprogrammer/xenia/pkg/apps"
	"github.com/transprogrammer/xenia/pkg/aspects"
	"github.com/transprogrammer/xenia/pkg/config"
	"github.com/transprogrammer/xenia/pkg/stacks"
)

func main() {
	app := apps.App

	cfg, err := config.Load()
	if err != nil {
		err = fmt.Errorf("load config: %w", err)
		log.Fatal(err)
	}

	core := stacks.NewCore(app, cfg)

	jumpBeat := core.JumpBeat()
	mongoBeats := core.MongoBeats()

	mongoDevBeat := mongoBeats.Development()
	mongoProdBeat := mongoBeats.Production()

	drums := [4]stacks.Drum{
		core,
		stacks.NewJump(app, cfg, jumpBeat),
		stacks.NewMongo(app, cfg, mongoDevBeat),
		stacks.NewMongo(app, cfg, mongoProdBeat),
	}

	for _, drum := range drums {
		aspects.AddTags(drum, cfg)
	}

	app.Synth()
}
