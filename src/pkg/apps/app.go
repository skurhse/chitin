package apps

import "github.com/hashicorp/terraform-cdk-go/cdktf"

func NewApp() cdktf.App {
	return cdktf.NewApp(nil)
}
