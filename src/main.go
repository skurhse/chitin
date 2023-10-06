package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/transprogrammer/xenia/pkg/cfg"
	"github.com/transprogrammer/xenia/pkg/stk"
)

func main() {
	app := cdktf.NewApp(nil)

	cfg, err := cfg.Load()
	if err != nil {
		err = fmt.Errorf("load config: %w", err)
		log.Fatal(err)
	}

	tokens := stk.NewTokens(cfg)

	core := stk.NewCore(app, cfg, tokens)

	jumpBeat := core.JumpBeat()
	mongoBeats := core.MongoBeats()
	mongoDevBeat := mongoBeats.Dev()
	mongoProdBeat := mongoBeats.Prod()
	mongoTokens := tokens.Mongo

	drums := [4]stk.Drum{
		core,
		stk.NewJump(app, cfg, jumpBeat, tokens.Jump),
		stk.NewMongo(app, cfg, mongoDevBeat, mongoTokens.Dev),
		stk.NewMongo(app, cfg, mongoProdBeat, mongoTokens.Prod),
	}

	for _, drum := range drums {
		// aspects.AddTags(drum, cfg)
		fmt.Printf(*drum.StackName())
	}

	app.Synth()
}
