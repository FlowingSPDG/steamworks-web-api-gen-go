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

func isValidType(t string) bool {
	switch t {
	case "int", "int32", "int64", "uint", "uint32", "uint64", "float32", "float64", "bool", "string":
		return true
	}
	// TODO: Support enum
	return false
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

func convertToString(baseType, value string) string {
	value = convertArg(value)
	switch baseType {
	case "int", "int32", "int64", "uint", "uint32", "uint64", "float32", "float64":
		return fmt.Sprintf("fmt.Sprintf(`%%v`, %s)", value)
	case "bool":
		return "strconv.FormatBool(" + value + ")"
	default:
		return value
	}
}

var (
	FuncMap = template.FuncMap{
		"isValidType":     isValidType,
		"convertArg":      convertArg,
		"convertToString": convertToString,
	}
)
