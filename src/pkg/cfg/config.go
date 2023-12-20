package cfg

import (
	"fmt"
	"os"
)

type Config interface {
	Name() string
	Regions() Regions
}

type Regions interface {
	Primary() *string
	Secondary() *string
}

type DefaultConfig struct {
	Name_    *string
	Regions_ DefaultRegions
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

func Load() (DefaultConfig, error) {

	cfg := DefaultConfig{
		Name_: new(string),
		Regions_: DefaultRegions{
			Primary_:   new(string),
			Secondary_: new(string),
		},
	}

	keyMap := map[*string]string{
		cfg.Name_:               EnvVarNames.Name,
		cfg.Regions_.Primary_:   EnvVarNames.Regions.Primary,
		cfg.Regions_.Secondary_: EnvVarNames.Regions.Secondary,
	}

	for k, v := range keyMap {
		var exists bool
		var val string

		*k, exists = os.LookupEnv(v)

		if !exists {
			err := fmt.Errorf("lookup env: %v nonexistent", v)
			return cfg, err
		}

		k = &val
	}

	return cfg, nil
}
