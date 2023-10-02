package cfg

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/jsii-runtime-go"
)

type Config interface {
	Name() *string
	Regions() Regions
}

type Regions interface {
	Primary() *string
	Secondary() *string
}

type DefaultConfig struct {
	Name_         *string
	Regions_      DefaultRegions
	WhitelistIPs_ *[]*string
}

type DefaultRegions struct {
	Primary_   *string
	Secondary_ *string
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

func Load() (cfg DefaultConfig, err error) {

	var whitelistIPsList string

	keyMap := map[*string]string{
		cfg.Name_:               EnvVarNames.Name,
		cfg.Regions_.Primary_:   EnvVarNames.Regions.Primary,
		cfg.Regions_.Secondary_: EnvVarNames.Regions.Secondary,
		&whitelistIPsList:       EnvVarNames.WhitelistIPs,
	}

	for k, v := range keyMap {
		var present bool
		*k, present = os.LookupEnv(v)
		if !present {
			err = fmt.Errorf("lookup env: %v not present", v)
			return
		}
	}

	whitelistIPVals := strings.Split(whitelistIPsList, ",")
	var whitelistIPs []*string
	for _, val := range whitelistIPVals {
		whitelistIPs = append(whitelistIPs, jsii.String(val))
	}

	cfg.WhitelistIPs_ = &whitelistIPs

	return
}
