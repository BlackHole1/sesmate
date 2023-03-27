package char

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Case string

const (
	CaseLower            Case = "lower"             // lower_case
	CaseUpper            Case = "upper"             // UPPER_CASE
	CaseCamel            Case = "camel"             // camelCase
	CasePascal           Case = "pascal"            // PascalCase
	CaseSnake            Case = "snake"             // snake_case
	CaseScreamingSnake   Case = "screaming_snake"   // SCREAMING_SNAKE_CASE
	CaseCapitalizedSnake Case = "capitalized_snake" // Capitalized_Snake_Case
)

func (c *Case) String() string {
	return string(*c)
}

func (c *Case) Set(val string) error {
	switch Case(val) {
	case CaseLower, CaseUpper, CaseCamel, CasePascal, CaseSnake, CaseScreamingSnake, CaseCapitalizedSnake:
		*c = Case(val)
		return nil
	default:
		return fmt.Errorf("invalid char case: %s", val)
	}
}

func (c *Case) Type() string {
	return "CharCase"
}

func Format(charCase Case, prefix, key string) string {
	val := prefix

	lower := cases.Lower(language.Und)
	upper := cases.Upper(language.Und)
	title := cases.Title(language.Und)

	switch charCase {
	case CaseLower:
		val += strings.ToLower(key)
	case CaseUpper:
		val += strings.ToUpper(key)
	case CaseCamel, CasePascal, CaseSnake, CaseScreamingSnake, CaseCapitalizedSnake:
		words := strings.FieldsFunc(key, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsDigit(r)
		})

		r := ""
		for index, word := range words {
			switch charCase {
			case CaseCamel:
				if index == 0 {
					r += lower.String(word)
				} else {
					r += title.String(word)
				}
			case CasePascal:
				r += title.String(word)
			case CaseSnake:
				r += lower.String(word)
				if index != len(words)-1 {
					r += "_"
				}
			case CaseScreamingSnake:
				r += upper.String(word)
				if index != len(words)-1 {
					r += "_"
				}
			case CaseCapitalizedSnake:
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
