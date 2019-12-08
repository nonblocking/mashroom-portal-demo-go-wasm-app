package main

import (
	"bytes"
	. "fmt"
	"text/template"
)

type Model struct {
	ResourcesBasePath string
	FirstName         string
}

var parsedTemplate, _ = template.New("htmlTemplate").Parse(`
	<div style="display:flex; padding:10px; align-items:center">
		<img src="{{.ResourcesBasePath}}/go_lang_logo.png" style="height:90px; margin:0 25px"/>
		<div>
			<h4>Go WebAssembly Demo App</h4>
			<p>
				Hello {{.FirstName}}!
			</p>
			<div>
				<button style="margin-right:10px">Send Ping</button>
				<span class="number-pings">
					Received pings: 0
				</span>
			</div>
		</div>
	</div>
`)

func createHtml(resourcesBasePath string, firstName string) string {
	model := Model{ResourcesBasePath: resourcesBasePath, FirstName: firstName}

	var buf bytes.Buffer
	err := parsedTemplate.Execute(&buf, model)
	if err != nil {
		Println("Error processing template!", err)
	}
	return buf.String()
}
