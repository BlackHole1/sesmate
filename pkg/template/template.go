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

func FindWithDir(dir string) ([]*SchemaBody, error) {
	abdPath, err := utils.ToAbs(dir, false)
	if err != nil {
		return nil, err
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	list := make([]string, 0, len(files))
	for _, dirEntry := range files {
		if !dirEntry.IsDir() && filepath.Ext(dirEntry.Name()) == ".json" {
			list = append(list, filepath.Join(abdPath, dirEntry.Name()))
		}
	}
	if len(list) == 0 {
		return nil, errors.New("no json file found")
	}

	result := make([]*SchemaBody, 0, len(list))

	for _, p := range list {
		t := validate(p)
		if t == nil {
			fmt.Printf("[sesmate]: skip %s, because template is invalid.\n", filepath.Base(p))
			continue
		}
		result = append(result, t)
	}
	if len(result) == 0 {
		return nil, errors.New("no template file found")
	}

	return result, err
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
