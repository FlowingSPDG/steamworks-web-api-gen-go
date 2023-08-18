package steamworkswebapigen

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/stoewer/go-strcase"
)

type TemplateInjection struct {
	AppVersion  string
	PackageName string

	Interfaces []Interface
}

func convertType(t string) string {
	switch t {
	case "int", "int32", "int64", "uint", "uint32", "uint64", "float32", "float64", "bool", "string":
		return t
	}
	return "any"
}

func camelCase(s string) string {
	return strcase.LowerCamelCase(s)
}

func convertArg(t string) string {
	if t == "type" {
		t = "t"
	}
	t = strcase.LowerCamelCase(t)

	if strings.HasSuffix(t, "[0]") {
		t = strings.TrimSuffix(t, "[0]")
	}

	return t
}

func convertToString(value string) string {
	value = convertArg(value)
	return fmt.Sprintf("fmt.Sprintf(`%%v`, input.%s)", value)
}

func getInputName(interfaceName string, method Method) string {
	return fmt.Sprintf("%s%s%dInput", interfaceName, method.Name, method.Version)
}

var (
	FuncMap = template.FuncMap{
		"convertType":     convertType,
		"convertArg":      convertArg,
		"convertToString": convertToString,
		"camelCase":       camelCase,
		"getInputName":    getInputName,
	}
)
