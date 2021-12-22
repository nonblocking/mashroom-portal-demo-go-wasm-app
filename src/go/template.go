package main

import (
	"bytes"
	. "fmt"
	"text/template"
)

type Model struct {
	ResourcesBasePath string
	Message string
	PingButtonLabel string
}

var parsedTemplate, _ = template.New("htmlTemplate").Parse(`
	<div style="display:flex; padding:10px; align-items:center">
		<img src="{{.ResourcesBasePath}}/go_lang_logo.png" style="height:90px; margin:0 25px"/>
		<div>
			<h4>Go WebAssembly Demo App</h4>
			<p>
				{{.Message}}!
			</p>
			<div>
				<button style="margin-right:10px">{{.PingButtonLabel}}</button>
				<span class="number-pings">
					Received pings: 0
				</span>
			</div>
		</div>
	</div>
`)

func createHtml(resourcesBasePath string, message string, pingButtonLabel string) string {
	model := Model{ResourcesBasePath: resourcesBasePath, Message: message, PingButtonLabel: pingButtonLabel}

	var buf bytes.Buffer
	err := parsedTemplate.Execute(&buf, model)
	if err != nil {
		Println("Error processing template!", err)
	}
	return buf.String()
}
