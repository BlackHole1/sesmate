package sync

import (
	"fmt"
	"log"

	awsSesV2 "github.com/aws/aws-sdk-go-v2/service/sesv2"

	"github.com/BlackHole1/sesmate/internal/sync/ses"
	"github.com/BlackHole1/sesmate/internal/sync/template"
)

type Context struct {
	client              *awsSesV2.Client
	remoteTemplateNames []string
	localTemplates      []*template.SchemaBody
	isRemove            bool
}

func New(ak, sk, endpoint, region, directory string, isRemove bool) *Context {
	if err := ses.Setup(ak, sk, endpoint, region); err != nil {
		log.Fatalln(err.Error())
	}

	localTemplates, err := template.FindWithDir(directory)
	if err != nil {
		log.Fatalln(err.Error())
	}

	localTemplateNames := make([]string, 0, len(localTemplates))
	for _, t := range localTemplates {
		localTemplateNames = append(localTemplateNames, t.TemplateName)
	}

	remoteTemplateNames, err := ses.ListTemplateName()
	if err != nil {
		log.Fatalln(err.Error())
	}

	return &Context{
		client:              ses.Client(),
		remoteTemplateNames: remoteTemplateNames,
		localTemplates:      localTemplates,
		isRemove:            isRemove,
	}
}

func (c *Context) Execute() error {
	if err := c.remove(); err != nil {
		return err
	}

	if err := c.update(); err != nil {
		return err
	}

	if err := c.create(); err != nil {
		return err
	}

	return nil
}

func (c *Context) remove() error {
	if !c.isRemove {
		return nil
	}

	for _, v := range c.remoteTemplateNames {
		t := template.FindWithName(c.localTemplates, v)
		if t == nil {
			if err := ses.RemoveTemplate(v); err != nil {
				return err
			}
			fmt.Printf("[sesmate]: Removed template: %s\n", v)
		}
	}

	return nil
}

// update If the templateName is the same for the local and remote files, we will update the remote file
// regardless of its content.
func (c *Context) update() error {
	for _, v := range c.remoteTemplateNames {
		t := template.FindWithName(c.localTemplates, v)

		if t != nil {
			if err := ses.UpdateTemplate(t); err != nil {
				return err
			}
			fmt.Printf("[sesmate]: Updated template: %s\n", t.TemplateName)
		}
	}

	return nil
}

func (c *Context) create() error {
	for _, v := range c.localTemplates {
		if !arrIn(c.remoteTemplateNames, v.TemplateName) {
			if err := ses.CreateTemplate(v); err != nil {
				return err
			}
			fmt.Printf("[sesmate]: Created template: %s\n", v.TemplateName)
		}
	}

	return nil
}

func arrIn(arr []string, target string) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}

	return false
}
