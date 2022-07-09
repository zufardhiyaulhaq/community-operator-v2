package message

const WEEKLY_TELEGRAM_TEMPLATE = `
*{{ .Name }}*
*{{ .Date }}*

{{ range .Articles }}• [{{ .Type }}] [{{ .Title }}]({{ .Url }})
{{ end }}

Tags: {{ range .Tags }}#{{ . }} {{ end }}
`
