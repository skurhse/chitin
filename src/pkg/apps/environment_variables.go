package apps

import (
	"fmt"
	"strings"
)

var EnvPrefix = strings.ToUpper(*AppName)
var RegionPrefix = fmt.Sprintf("%s_REGION", EnvPrefix)

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
	Name: fmt.Sprintf("%s_NAME", EnvPrefix),
	Regions: RegionEnvVarNamesIndex{
		Primary:   fmt.Sprintf("%s_PRIMARY", RegionPrefix),
		Secondary: fmt.Sprintf("%s_SECONDARY", RegionPrefix),
	},
	WhitelistIPs: fmt.Sprintf("%s_WHITELIST_IPS", EnvPrefix),
}

type EnvVarName
