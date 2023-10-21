package mod

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/transprogrammer/xenia/generated/naming"
)

type NamingConfig interface {
	Tokens() []*string
}

func NewNaming(scope constructs.Construct, tokens []string) naming.Naming {
	var prefix []*string
	var id = *Ids.Naming
	for _, token := range tokens {
		prefix = append(prefix, jsii.String(token))
		id = fmt.Sprintf("%s_%s", id, token)
	}

	input := naming.NamingConfig{
		Prefix:               &prefix,
		UniqueIncludeNumbers: jsii.Bool(false),
		Suffix:               &[]*string{},
	}

	return naming.NewNaming(scope, &id, &input)
}
