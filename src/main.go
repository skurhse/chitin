package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/skurhse/chitin/pkg/cfg"
	"github.com/skurhse/chitin/pkg/cre"
)

func main() {
	app := cdktf.NewApp(nil)

	appCfg, err := cfg.Load()
	if err != nil {
		err = fmt.Errorf("load config: %w", err)
		log.Fatal(err)
	}

	tokens := cfg.Tokens

	cre.NewCore(app, appCfg, tokens.Core)

	app.Synth()
}
