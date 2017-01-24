``

	{{with .flash_success}}<div class="alert alert-success">{{.}}</div>{{end}}
	{{with .flash_error}}<div class="alert alert-danger">{{.}}</div>{{end}}
	{{template "yield" .}}
	{{template "authboss" .}}

{{define "pagetitle"}}{{end}}
{{define "yield"}}{{end}}
{{define "authboss"}}{{end}}
