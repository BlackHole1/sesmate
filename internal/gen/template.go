package gen

const templateStr = `package {{ .PackageName }}
{{ range .Items }}
const {{ .Variable }} = "{{ .Constant }}"

{{ if eq (len .SESData) 0 -}}
type {{ .Variable }}Data struct {}
{{ else -}}
type {{ .Variable }}Data struct {
{{ range .SESData -}}
	{{ .Key }} string ` + "`" + `json:"{{ .JsonKey }}"` + "`" + `
{{ end -}}
}
{{ end -}}
{{ end -}}`
