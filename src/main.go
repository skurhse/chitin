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
	jump := stk.NewJump(app, cfg, core.JumpBeat(), tokens.Jump)

	mongoDev := stk.NewMongo(app, cfg, core.MongoBeats().Dev(), tokens.Mongo.Dev)
	mongoProd := stk.NewMongo(app, cfg, core.MongoBeats().Prod(), tokens.Mongo.Prod)

	cluster := stk.NewCluster(app, cfg, core.ClusterBeat(), jump.ClusterBeat(), tokens.Cluster)

	drums := [5]stk.Drum{core, jump, mongoDev, mongoProd, cluster}

	for _, drum := range drums {
		// aspects.AddTags(drum, cfg)
		fmt.Printf(*drum.StackName())
	}

	app.Synth()
}
