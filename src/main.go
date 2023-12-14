package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/skurhse/xen/pkg/cfg"
	"github.com/skurhse/xen/pkg/cre"
)

func main() {
	app := cdktf.NewApp(nil)

	config, err := cfg.Load()
	if err != nil {
		err = fmt.Errorf("load config: %w", err)
		log.Fatal(err)
	}

	tokenSets := cfg.NewTokenSets(config)
	tokens := cfg.Tokens

	cre.NewCore(app, config, tokenSets, tokens.Core)

	app.Synth()
}
