package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/skurhse/xen/pkg/cfg"
)

func main() {
	app := cdktf.NewApp(nil)

	cfg, err := cfg.Load()
	if err != nil {
		err = fmt.Errorf("load config: %w", err)
		log.Fatal(err)
	}

	tokenSets := cfg.NewTokenSets(cfg)
	tokens := cfg.Tokens

	sng.NewCore(app, cfg, tokenSets, tokens.Core)

	app.Synth()
}
