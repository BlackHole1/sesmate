package gen

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	sestemplate "github.com/BlackHole1/sesmate/pkg/template"
	"github.com/BlackHole1/sesmate/pkg/utils"
)

type CharCase string

const (
	CharCaseLower            CharCase = "lower"             // lower_case
	CharCaseUpper            CharCase = "upper"             // UPPER_CASE
	CharCaseCamel            CharCase = "camel"             // camelCase
	CharCasePascal           CharCase = "pascal"            // PascalCase
	CharCaseSnake            CharCase = "snake"             // snake_case
	CharCaseScreamingSnake   CharCase = "screaming_snake"   // SCREAMING_SNAKE_CASE
	CharCaseCapitalizedSnake CharCase = "capitalized_snake" // Capitalized_Snake_Case
)

func (c *CharCase) String() string {
	return string(*c)
}

func (c *CharCase) Set(val string) error {
	switch CharCase(val) {
	case CharCaseLower, CharCaseUpper, CharCaseCamel, CharCasePascal, CharCaseSnake, CharCaseScreamingSnake, CharCaseCapitalizedSnake:
		*c = CharCase(val)
		return nil
	default:
		return fmt.Errorf("invalid char case: %s", val)
	}
}

func (c *CharCase) Type() string {
	return "CharCase"
}

const templateStr = `package {{ .PackageName }}
{{ range .Names }}
const {{ .Variable }} = "{{ .Constant }}"
{{ end }}`

type TemplateData struct {
	PackageName string
	Names       []Names
}

type Names struct {
	Variable string
	Constant string
}

type Context struct {
	names       []Names
	outputFile  string
	packageName string
}

func New(dir, output, filename, packageName, prefix string, charCase CharCase) *Context {
	localTemplates, err := sestemplate.FindWithDir(dir)
	if err != nil {
		log.Fatalln(err.Error())
	}

	names := make([]Names, 0, len(localTemplates))
	for _, t := range localTemplates {
		names = append(names, Names{
			Variable: format(charCase, prefix, t.TemplateName),
			Constant: t.TemplateName,
		})
	}

	outputPath, err := utils.ToAbs(output, true)
	if err != nil {
		log.Fatalln(err.Error())
	}
	outputFile := filepath.Join(outputPath, filename+".go")

	return &Context{
		names:       names,
		outputFile:  outputFile,
		packageName: packageName,
	}
}

func (c *Context) Execute() error {
	tmplObj, err := template.New("template").Parse(templateStr)
	if err != nil {
		return err
	}

	outputFile, err := os.Create(c.outputFile)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	err = tmplObj.Execute(outputFile, TemplateData{
		PackageName: c.packageName,
		Names:       c.names,
	})
	if err != nil {
		return err
	}

	return nil
}

func format(charCase CharCase, prefix, key string) string {
	val := prefix

	lower := cases.Lower(language.Und)
	upper := cases.Upper(language.Und)
	title := cases.Title(language.Und)

	switch charCase {
	case CharCaseLower:
		val += strings.ToLower(key)
	case CharCaseUpper:
		val += strings.ToUpper(key)
	case CharCaseCamel, CharCasePascal, CharCaseSnake, CharCaseScreamingSnake, CharCaseCapitalizedSnake:
		words := strings.FieldsFunc(key, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsDigit(r)
		})

		r := ""
		for index, word := range words {
			switch charCase {
			case CharCaseCamel:
				if index == 0 {
					r += lower.String(word)
				} else {
					r += title.String(word)
				}
			case CharCasePascal:
				r += title.String(word)
			case CharCaseSnake:
				r += lower.String(word)
				if index != len(words)-1 {
					r += "_"
				}
			case CharCaseScreamingSnake:
				r += upper.String(word)
				if index != len(words)-1 {
					r += "_"
				}
			case CharCaseCapitalizedSnake:
				r += title.String(word)
				if index != len(words)-1 {
					r += "_"
				}
			}
		}

		val += r
	}

	return val
}
