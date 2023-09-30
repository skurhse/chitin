package apps

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

var AppConfig = &cdktf.AppConfig{
	Outdir: jsii.String("./out"),
}

var App = cdktf.NewApp(AppConfig)
