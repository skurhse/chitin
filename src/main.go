package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/skurhse/xen/pkg/cfg"
	"github.com/skurhse/xen/pkg/sng"
)

func main() {
	app := cdktf.NewApp(nil)

	cfg, err := cfg.Load()
	if err != nil {
		err = fmt.Errorf("load config: %w", err)
		log.Fatal(err)
	}

	tokenSets := sng.NewTokenSets(cfg)
	tokens := sng.Tokens

	sng.NewCore(app, cfg, tokenSets, tokens.Core)

	app.Synth()
}
