package config

import (
	"fmt"
	"strings"

	"github.com/transprogrammer/xenia/pkg/apps"
)

var EnvVarPrefix = strings.ToUpper(*apps.AppName)

var RegionEnvVarPrefix = fmt.Sprintf("%s_REGION", EnvVarPrefix)

type EnvVarNamesIndex struct {
	Name         string
	Regions      RegionEnvVarNamesIndex
	WhitelistIPs string
}

type RegionEnvVarNamesIndex struct {
	Primary   string
	Secondary string
}

var EnvVarNames = EnvVarNamesIndex{
	Name: fmt.Sprintf("%s_NAME", EnvVarPrefix),
	Regions: RegionEnvVarNamesIndex{
		Primary:   fmt.Sprintf("%s_PRIMARY", RegionEnvVarPrefix),
		Secondary: fmt.Sprintf("%s_SECONDARY", RegionEnvVarPrefix),
	},
	WhitelistIPs: fmt.Sprintf("%s_WHITELIST_IPS", EnvVarPrefix),
}
