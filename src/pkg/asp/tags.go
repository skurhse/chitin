package asp

import (
	"maps"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/skurhse/chitin/pkg/cfg"
	"github.com/skurhse/chitin/pkg/stk"
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

func AddTags(container stk.Drum, cfg cfg.Config) {
	stack := container.Stack()
	stackName := container.StackName()
	projectName := cfg.Name()

	projectEntry := &map[string]*string{ProjectTagKey: projectName}
	stackEntry := &map[string]*string{StackTagKey: stackName}
	entries := [2]*map[string]*string{projectEntry, stackEntry}

	asp := cdktf.Aspects_Of(stack)

	for _, entry := range entries {
		aspect := NewTagsAddingAspect(entry)
		asp.Add(aspect)
	}
}
