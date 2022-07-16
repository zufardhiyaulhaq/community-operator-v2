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

const MEETUP_TELEGRAM_TEMPLATE = `
*{{ .Name }}*

Hi Everyone!

Let's gather again at our meetup sponsored by {{ range .Sponsors }}{{ .Name }}, {{ end }}!

Our Speaker:
{{ range .Speakers }}•  {{ .Name }}, {{ .Position }} @ {{ .Company }} sharing with us on "{{ .Title }}"
{{ end }}

The meetup will be held on {{ .Date }}, at {{ .Place }}. Doors will be open at {{ .Time }}. 

RSVP your seat at {{ .RegistrationUrl }} . See you!
Tags: {{ range .Tags }}#{{ . }} {{ end }}
`

const WEEKLY_TWITTER_TEMPLATE = `
{{ .Name }}
{{ .Date }}

{{ range .Articles }}• [{{ .Type }}] {{ .Title }}: {{ .Url }}
{{ end }}

Tags: {{ range .Tags }}#{{ . }} {{ end }}
`

const ANNOUNCEMENT_TWITTER_TEMPLATE = `
{{ .Subject }}

Hi Everyone!

{{ .Body }}

Tags: {{ range .Tags }}#{{ . }} {{ end }}
`

const MEETUP_TWITTER_TEMPLATE = `
{{ .Name }}

Hi Everyone!

Let's gather again at our meetup sponsored by {{ range .Sponsors }}{{ .Name }}, {{ end }}!

Our Speakers:
{{ range .Speakers }}•  {{ .Name }}, {{ .Position }} @ {{ .Company }} sharing with us on "{{ .Title }}"
{{ end }}

The meetup will be held on {{ .Date }}, at {{ .Place }}. Doors will be open at {{ .Time }}. 

RSVP your seat at {{ .RegistrationUrl }} . See you!
Tags: {{ range .Tags }}#{{ . }} {{ end }}
`
