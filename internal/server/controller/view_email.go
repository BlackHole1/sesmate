package controller

import (
	_ "embed"
	"encoding/json"
	"html/template"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/BlackHole1/sesmate/internal/server/svc"
)

//go:embed view_email.tmpl
var viewEmailTmpl string

type ViewEmailData struct {
	Email []string
}

func ViewEmail(c *gin.Context) {
	tmpl, err := template.New("view_email").Parse(viewEmailTmpl)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	list, err := svc.Email.AllRecord()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	result := make([]string, 0, len(list))
	for _, item := range list {
		var jsonObj map[string]any
		err = json.Unmarshal([]byte(item.RawData), &jsonObj)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		formattedJson, err := json.MarshalIndent(jsonObj, "", "  ")
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		text := string(formattedJson)
		text = strings.Replace(text, "\\u003c", "<", -1)
		text = strings.Replace(text, "\\u003e", ">", -1)
		text = strings.Replace(text, "\\u0026", "&", -1)

		result = append(result, text)
	}

	tmpl.Execute(c.Writer, result)
}
