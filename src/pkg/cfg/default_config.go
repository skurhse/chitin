package cfg

import (
	"fmt"
	"os"
)

type DefaultConfig struct {
	Name_    string
	Regions_ DefaultRegions
}

type DefaultRegions struct {
	Primary_   string
	Secondary_ string
}

func (c DefaultConfig) Name() string {
	return c.Name_
}

func (c DefaultConfig) Regions() Regions {
	return c.Regions_
}

func (r DefaultRegions) Primary() string {
	return r.Primary_
}

func (r DefaultRegions) Secondary() string {
	return r.Secondary_
}

func Load() (DefaultConfig, error) {

	cfg := DefaultConfig{}

	keyMap := map[*string]string{
		&cfg.Name_:               EnvVarNames.Name,
		&cfg.Regions_.Primary_:   EnvVarNames.Regions.Primary,
		&cfg.Regions_.Secondary_: EnvVarNames.Regions.Secondary,
	}

	for k, v := range keyMap {
		var exists bool

		*k, exists = os.LookupEnv(v)

		if !exists {
			err := fmt.Errorf("lookup env: %v nonexistent", v)
			return cfg, err
		}
	}

	return cfg, nil
}
