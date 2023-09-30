package stacks

import "github.com/hashicorp/terraform-cdk-go/cdktf"

type SecretVariablesIndex struct {
	Name    cdktf.TerraformVariable
	Regions RegionSecretVariablesIndex
}

type RegionSecretVariablesIndex struct {
	Primary   cdktf.TerraformVariable
	Secondary cdktf.TerraformVariable
}

func NewSecretVariables(stack cdktf.TerraformStack) SecretVariablesIndex {
	cdktf.NewTerraformVariable(stack, jsii.String("imageId"), &cdktf.TerraformVariableConfig{
    Type:        jsii.String("string"),
    Default:     jsii.String("ami-abcde123"),
    Description: jsii.String("What AMI to use to create an instance"),
})
