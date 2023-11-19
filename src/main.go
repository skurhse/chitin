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

	tokenSets := stk.NewTokenSets(cfg)
	tokens := stk.Tokens

	core := stk.NewCore(app, cfg, tokenSets, tokens.Core)
	jump := stk.NewJump(app, cfg, core.JumpBeat(), tokenSets.Jump)
	postgres := stk.NewPostgres(app, cfg, core.PostgresBeat().Dev(), tokens.Postgres)

	app.Synth()
}
