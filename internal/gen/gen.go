package gen

import (
	"bytes"
	goformat "go/format"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/BlackHole1/sesmate/pkg/char"
	sestemplate "github.com/BlackHole1/sesmate/pkg/template"
	"github.com/BlackHole1/sesmate/pkg/utils"
)

type TemplateData struct {
	PackageName string
	Items       []*SESMemberItem
}

type SESMemberItem struct {
	Variable string
	Constant string
	SESData  []*SESDatum
}

type Context struct {
	items       []*SESMemberItem
	outputFile  string
	packageName string
}

func New(dir, output, filename, packageName, prefix string, charCase char.Case) *Context {
	localTemplates, err := sestemplate.FindDir(dir)
	if err != nil {
		log.Fatalln(err.Error())
	}

	items := make([]*SESMemberItem, 0, len(localTemplates))
	for _, t := range localTemplates {
		items = append(items, &SESMemberItem{
			Variable: char.Format(charCase, prefix, t.TemplateName),
			Constant: t.TemplateName,
			SESData:  getSESData(t),
		})
	}

	outputPath, err := utils.ToAbs(output, true)
	if err != nil {
		log.Fatalln(err.Error())
	}
	outputFile := filepath.Join(outputPath, filename+".go")

	return &Context{
		items:       items,
		outputFile:  outputFile,
		packageName: packageName,
	}
}

func (c *Context) Execute() error {
	tmplObj, err := template.New("template").Parse(templateStr)
	if err != nil {
		return err
	}

	var tplOutput bytes.Buffer
	err = tmplObj.Execute(&tplOutput, TemplateData{
		PackageName: c.packageName,
		Items:       c.items,
	})

	if err != nil {
		return err
	}

	source, err := goformat.Source(tplOutput.Bytes())
	if err != nil {
		return err
	}

	outputFile, err := os.Create(c.outputFile)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	if _, err = outputFile.Write(source); err != nil {
		return err
	}

	return nil
}
