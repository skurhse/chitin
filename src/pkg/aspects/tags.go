package aspects

import (
	"maps"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/transprogrammer/xenia/pkg/config"
	"github.com/transprogrammer/xenia/pkg/stacks"
)

const ProjectTagKey = "project"
const StackTagKey = "stack"

type Taggable interface {
	TagsInput() *map[string]*string
	SetTags(val *map[string]*string)
}

type TagsAddingAspect struct {
	Tags *map[string]*string
}

func (taa TagsAddingAspect) Visit(node constructs.IConstruct) {
	if taggable, ok := node.(Taggable); ok {
		existing := *taggable.TagsInput()
		tags := *taa.Tags
		maps.Copy(existing, tags) // requires Go 1.18
		taggable.SetTags(&existing)
	}
}

func NewTagsAddingAspect(tags *map[string]*string) *TagsAddingAspect {
	return &TagsAddingAspect{Tags: tags}
}

func AddTags(container stacks.StackContainer, config config.AppConfig) {
	stack := *container.Stack()
	stackName := container.StackName()
	projectName := config.Name()

	projectEntry := &map[string]*string{ProjectTagKey: projectName}
	stackEntry := &map[string]*string{StackTagKey: stackName}
	entries := [2]*map[string]*string{projectEntry, stackEntry}

	aspects := cdktf.Aspects_Of(stack)

	for _, entry := range entries {
		aspect := NewTagsAddingAspect(entry)
		aspects.Add(aspect)
	}
}
