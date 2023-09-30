package apps

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/aws/jsii-runtime-go"
)

const configPath = "app.json"

var AppName = jsii.String("xenia")

type Config interface {
	Name() *string
	Regions() Regions
	SSHSourceAddressPrefixes() *[]*string
	MongoEnvironments() MongoEnvironments
}

type Regions interface {
	Primary() *string
	Secondary() *string
}

type DefaultConfig struct {
	Name_                     *string        `json:"name"`
	SubscriptionId_           *string        `json:"subscriptionId"`
	Regions_                  DefaultRegions `json:"regions"`
	SSHSourceAddressPrefixes_ *[]*string     `json:"sshSourceAddressPrefixes"`
	MongoEnvironments_        DefaultMongoEnvironments
}

type DefaultRegions struct {
	Primary_   *string `json:"primary"`
	Secondary_ *string `json:"secondary"`
}

func (c DefaultConfig) SubscriptionId() *string {
	return c.SubscriptionId_
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

func (c DefaultConfig) SSHSourceAddressPrefixes() *[]*string {
	return c.SSHSourceAddressPrefixes_
}

func (c DefaultConfig) MongoEnvironments() MongoEnvironments {
	return c.MongoEnvironments_
}

func (e DefaultMongoEnvironments) Development() MongoEnvironment {
	return e.Development_
}

func (e DefaultMongoEnvironments) Production() MongoEnvironment {
	return e.Production_
}

func LoadConfig() (cfg DefaultConfig, err error) {

	file, err := os.Open(configPath)
	if err != nil {
		return cfg, fmt.Errorf("open %v: %w", configPath, err)
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		panic(err)
	}

	return cfg, nil
}
