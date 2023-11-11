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
	//jump := stk.NewJump(app, cfg, core.JumpBeat(), tokens.Jump)

	postgresDev := stk.NewPostgres(app, cfg, core.PostgresBeats().Dev(), tokens.Postgres.Dev)
	// postgresProd := stk.NewPostgres(app, cfg, core.PostgresBeats().Prod(), tokens.Postgres.Prod)

	// cluster := stk.NewCluster(app, cfg, core.ClusterBeat(), jump.ClusterBeat(), tokens.Cluster)

	// drums := [5]stk.Drum{core, jump, postgresDev, postgresProd, cluster}
	drums := [5]stk.Drum{core, postgresDev}

	for _, drum := range drums {
		// asp.AddTags(drum, cfg)
		fmt.Printf(*drum.StackName())
	}

	app.Synth()
}
