package stacks

import (
	"fmt"
	"os"
	"strings"
)

type Config interface {
	Name() *string
	Regions() Regions
	WhitelistIPs() *[]*string
}

type Regions interface {
	Primary() *string
	Secondary() *string
}

type DefaultConfig struct {
	Name_         *string        `json:"name"`
	Regions_      DefaultRegions `json:"regions"`
	WhitelistIPs_ *[]*string     `json:"sshSourceAddressPrefixes"`
}

type DefaultRegions struct {
	Primary_   *string `json:"primary"`
	Secondary_ *string `json:"secondary"`
}

func (c DefaultConfig) Name() *string {
	return c.Name_
}

func (c DefaultConfig) Regions() Regions {
	return c.Regions_
}

func (r DefaultRegions) Primary() *string {
	return r.Primary_
}

func (r DefaultRegions) Secondary() *string {
	return r.Secondary_
}

func (c DefaultConfig) WhitelistIPs() *[]*string {
	return c.WhitelistIPs_
}

func LoadConfig() (cfg DefaultConfig, err error) {

	var whitelistIPsList *string

	keyMap := map[*string]string{
		cfg.Name_:               EnvVarNames.Name,
		cfg.Regions_.Primary_:   EnvVarNames.Regions.Primary,
		cfg.Regions_.Secondary_: EnvVarNames.Regions.Secondary,
		whitelistIPsList:        EnvVarNames.WhitelistIPs,
	}

	for k, v := range keyMap {
		var present bool
		*k, present = os.LookupEnv(v)
		if !present {
			err = fmt.Errorf("lookup env: %v not present", v)
			return
		}
	}

	whitelistIPVals := strings.Split(*whitelistIPsList, ",")
	var whitelistIPs []*string
	for i, _ := range whitelistIPVals {
		val := whitelistIPVals[i]
		whitelistIPs = append(whitelistIPs, &val)
	}

	cfg.WhitelistIPs_ = &whitelistIPs

	return
}
