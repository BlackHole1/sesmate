package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/BlackHole1/sesmate/pkg/template"
	"github.com/BlackHole1/sesmate/pkg/utils"
)

const kTemplatePath = "sesmate-template-path"

func Template(dir string) func(c *gin.Context) {
	p, err := utils.ToAbs(dir, true)
	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		c.Set(kTemplatePath, p)
		c.Next()
	}
}

func MustTemplateDir(c *gin.Context) string {
	return c.MustGet(kTemplatePath).(string)
}

func HasTemplate(c *gin.Context, name string) bool {
	p := MustTemplateDir(c)

	localTemplates, err := template.FindDirWithoutWarn(p)
	if err != nil {
		return false
	}

	for _, t := range localTemplates {
		if t.TemplateName == name {
			return true
		}
	}

	return false
}
