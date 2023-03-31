package template

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/BlackHole1/sesmate/pkg/utils"
)

type SchemaBody struct {
	HtmlPart     *string `json:"HtmlPart"`
	SubjectPart  *string `json:"SubjectPart"`
	TemplateName string  `json:"TemplateName"`
	TextPart     *string `json:"TextPart"`
}

type Schema struct {
	Template SchemaBody `json:"Template"`
}

func findDir(dir string, showWarn bool) ([]*SchemaBody, error) {
	abdPath, err := utils.ToAbs(dir, false)
	if err != nil {
		return nil, err
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	list := make([]*SchemaBody, 0, len(files))
	for _, dirEntry := range files {
		if !dirEntry.IsDir() && filepath.Ext(dirEntry.Name()) == ".json" {
			p := filepath.Join(abdPath, dirEntry.Name())
			t := validate(p)
			if t == nil {
				if showWarn {
					fmt.Printf("[sesmate]: skip %s, because template is invalid.\n", filepath.Base(p))
				}
				continue
			}

			if info, err := os.Stat(p); err != nil {
				if showWarn {
					fmt.Printf("[sesmate]: skip %s, because %s.\n", filepath.Base(p), err.Error())
				}
				continue
			} else if info.Size() > 500*1024 {
				if showWarn {
					fmt.Printf("[sesmate]: skip %s, because template size exceeds the limit of 500 KB.\n", filepath.Base(p))
				}
				continue
			}

			list = append(list, t)
		}
	}
	if len(list) == 0 {
		return nil, errors.New("no template file found")
	}
	if len(list) > 10000 {
		return nil, errors.New("the number of templates exceeds the limit of 10000")
	}

	return list, err
}

func FindDir(dir string) ([]*SchemaBody, error) {
	return findDir(dir, true)
}

func FindDirWithoutWarn(dir string) ([]*SchemaBody, error) {
	return findDir(dir, false)
}

func FindWithName(body []*SchemaBody, templateName string) *SchemaBody {
	for _, v := range body {
		if v.TemplateName == templateName {
			return v
		}
	}

	return nil
}

func validate(p string) *SchemaBody {
	file, err := os.Open(p)
	if err != nil {
		return nil
	}
	defer file.Close()

	var template Schema
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&template)
	if err != nil {
		return nil
	}

	if template.Template.TemplateName == "" {
		return nil
	}

	return &template.Template
}
