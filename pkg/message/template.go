package message

const WEEKLY_TELEGRAM_TEMPLATE = `
*{{ .Name }}*
*{{ .Date }}*

{{ range .Articles }}• [{{ .Type }}] [{{ .Title }}]({{ .Url }})
{{ end }}

Tags: {{ range .Tags }}#{{ . }} {{ end }}
`

const ANNOUNCEMENT_TELEGRAM_TEMPLATE = `
*{{ .Subject }}*

Hi Everyone!

{{ .Body }}

Tags: {{ range .Tags }}#{{ . }} {{ end }}
`
