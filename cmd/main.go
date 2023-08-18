package main

import (
	"bytes"
	"go/format"
	"html/template"
	"log"
	"os"

	_ "embed"

	steamworkswebapigen "github.com/FlowingSPDG/steamworks-web-api-gen-go"
)

const (
	packageName = "steamworks"
	appVersion  = "v0.1.0"
)

var (
	//go:embed templates/interface.tmpl
	templateInterface string

	//go:embed templates/response.tmpl
	templateResponse string

	//go:embed templates/request.tmpl
	templateRequest string
)

func main() {
	// parse
	key := os.Getenv("STEAM_WEB_API_KEY")

	// check flags
	if key == "" {
		panic("Steam Web API key is required")
	}

	// Get interfaces
	resp, err := steamworkswebapigen.GetSupportedAPIList(key)
	if err != nil {
		panic(err)
	}

	// set template
	injection := steamworkswebapigen.TemplateInjection{
		AppVersion:  appVersion,
		PackageName: packageName,
		Interfaces:  resp.Apilist.Interfaces,
	}

	// Remove generated files
	os.RemoveAll("./generated/")

	// create directory
	if err := os.MkdirAll("./generated/", 0777); err != nil {
		panic(err)
	}

	// execute
	templates := []struct {
		template   string
		targetFile string
	}{
		{templateInterface, "./generated/interface.go"},
		{templateResponse, "./generated/response.go"},
		{templateRequest, "./generated/request.go"},
	}
	for _, tmpl := range templates {
		// create buffer
		buf := &bytes.Buffer{}

		t := template.Must(template.New("get_api_list").Funcs(steamworkswebapigen.FuncMap).Parse(tmpl.template))
		if err := t.Execute(buf, injection); err != nil {
			log.Fatalf("Failed to execute template %s", err)
		}

		by := buf.Bytes()
		b, err := format.Source(by)
		if err != nil {
			log.Fatalf("Failed to source %v %s", err, by)
		}

		// create file
		if err := os.WriteFile(tmpl.targetFile, b, 0777); err != nil {
			log.Fatalf("Failed to write file %s", err)
		}
	}

	// TODO: gopls/goimports
}
