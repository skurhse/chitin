package modules

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/apps"
)

type NamingConfig interface {
	Tokens() []*string
}

var NamingPrefix = &[]*string{apps.Name}

func NewNaming(scope constructs.Construct, config NamingConfig, suffix *[]*string) naming.Naming {
	tokens := config.Tokens()

	input := &naming.NamingConfig{
		Prefix:               &tokens,
		UniqueIncludeNumbers: jsii.Bool(false),
		Suffix:               &[]*string{},
	}

	return naming.NewNaming(scope, Ids.Naming, input)
}
