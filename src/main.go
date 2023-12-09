package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/skurhse/xen/pkg/cfg"
	"github.com/skurhse/xen/pkg/sng"

ck
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
	stk.NewJump(app, cfg, core.JumpBeat(), tokenSets.Jump)
	stk.NewPostgres(app, cfg, core.PostgresBeat(), tokenSets.Postgres)

	app.Synth()
}
