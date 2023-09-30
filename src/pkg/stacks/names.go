package stacks

import (
	"fmt"

	"github.com/aws/jsii-runtime-go"
)

type StackNamesIndex struct {
	Core  *string
	Mongo MongoStackNamesIndex
	Jump  *string
}

type MongoStackNamesIndex struct {
	Development *string
	Production  *string
}

type StackTokensIndex struct {
	Core  *[]*string
	Mongo MongoStackTokensIndex
	Jump  *[]*string
}

type MongoStackTokensIndex struct {
	Development *[]*string
	Production  *[]*string
}

var StackTokens = StackTokensIndex{
	Core: &[]*string{jsii.String("core")},
	Mongo: MongoStackTokensIndex{
		Development: &[]*string{jsii.String("mongo"), jsii.String("dev")},
		Production:  &[]*string{jsii.String("mongo"), jsii.String("prod")},
	},
	Jump: &[]*string{jsii.String("jump")},
}

var StackNames = StackNamesIndex{
	Core: (*StackTokens.Core)[0],
	Mongo: MongoStackNamesIndex{
		Development: jsii.String(fmt.Sprintf("%s-%s",
			(*StackTokens.Mongo.Development)[0], (*StackTokens.Mongo.Development)[1])),
		Production: jsii.String(fmt.Sprintf("%s-%s",
			(*StackTokens.Mongo.Production)[0], (*StackTokens.Mongo.Development)[1])),
	},
	Jump: StackTokens.Jump[0],
}
