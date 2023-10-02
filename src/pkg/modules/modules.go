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

func NewNaming(scope constructs.Construct, tokens []string) naming.Naming {
	var prefix []*string
	for _, token := range tokens {
		prefix = append(prefix, jsii.String(token))
	}

	input := &naming.NamingConfig{
		Prefix:               &prefix,
		UniqueIncludeNumbers: jsii.Bool(false),
		Suffix:               &[]*string{},
	}

	return naming.NewNaming(scope, Ids.Naming, input)
}
