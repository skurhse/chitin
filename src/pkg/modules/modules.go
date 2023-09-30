package modules

import (
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/apps"
)

type NamingConfig interface {
	config.Config
}

func NewNaming(scope constructs.Contstruct, config NamingConfig, suffix *[]*string) *naming.Naming {

	id := ModuleIds.Naming

	prefix := &[]*string{apps.AppName}

	input := naming.NamingConfig{
		Prefix:               prefix,
		UniqueIncludeNumbers: jsii.Bool(false),
		Suffix:               suffix,
	}

	naming := naming.NewNaming(scope, id, &input)

	return &naming
}
