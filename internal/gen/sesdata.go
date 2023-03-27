package gen

import (
	"fmt"
	"regexp"

	"github.com/BlackHole1/sesmate/pkg/char"
	sestemplate "github.com/BlackHole1/sesmate/pkg/template"
)

type SESDatum struct {
	Key     string
	JsonKey string
}

var sesTemplateDataReg = regexp.MustCompile(`{{\s*([a-zA-Z_][a-zA-Z0-9_]*)\s*}}`)

func getSESData(t *sestemplate.SchemaBody) []*SESDatum {
	var list []*SESDatum
	keyCount := make(map[string]int)

	sesDataWithPart(&list, &keyCount, t.HtmlPart)
	sesDataWithPart(&list, &keyCount, t.TextPart)

	return list
}

func sesDataWithPart(list *[]*SESDatum, keyCount *map[string]int, str *string) {
	if str == nil || *str == "" {
		return
	}

	matches := sesTemplateDataReg.FindAllStringSubmatch(*str, -1)
	for _, match := range matches {
		key := char.Format(char.CasePascal, "", match[1])

		if inJsonKey(*list, match[1]) {
			continue
		}

		var datum *SESDatum

		if inKey(*list, key) {
			// start from 2
			if (*keyCount)[key] == 0 {
				(*keyCount)[key] = 1
			}
			(*keyCount)[key]++

			datum = &SESDatum{
				Key:     fmt.Sprintf("%s%d", key, (*keyCount)[key]),
				JsonKey: match[1],
			}
		} else {
			datum = &SESDatum{
				Key:     key,
				JsonKey: match[1],
			}
		}

		*list = append(*list, datum)
	}
}

func inJsonKey(list []*SESDatum, key string) bool {
	for _, v := range list {
		if v.JsonKey == key {
			return true
		}
	}

	return false
}

func inKey(list []*SESDatum, key string) bool {
	for _, v := range list {
		if v.Key == key {
			return true
		}
	}

	return false
}
