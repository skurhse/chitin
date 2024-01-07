package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/skurhse/chitin/pkg/cfg"
	"github.com/skurhse/chitin/pkg/sng/bck"
	"github.com/skurhse/chitin/pkg/sng/cre"
)

func main() {
	app := cdktf.NewApp(nil)

	env, err := cfg.LoadEnvironment()
	if err != nil {
		err = fmt.Errorf("loading environment: %w", err)
		log.Fatal(err)
	}

	tokens := cfg.Tokens

	bck.NewBackend(app, env, tokens.Backend)
	cre.NewCore(app, env, tokens.Core)

	app.Synth()
}
