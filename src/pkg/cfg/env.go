package cfg

import (
	"fmt"
	"os"
)

type Environment struct {
	Name_    string
	Regions_ EnvironmentRegions
}

type EnvironmentRegions struct {
	Primary_   string
	Secondary_ string
}

func (c Environment) Name() string {
	return c.Name_
}

func (c Environment) Regions() Regions {
	return c.Regions_
}

func (r EnvironmentRegions) Primary() string {
	return r.Primary_
}

func (r EnvironmentRegions) Secondary() string {
	return r.Secondary_
}

func LoadEnvironment() (Environment, error) {

	cfg := Environment{}

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
