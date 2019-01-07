package provider

import (
	"text/template"

	"github.com/Masterminds/sprig"
)

// nolint: gochecknoglobals
var funcMap template.FuncMap

// nolint: gochecknoinits
func init() {
	funcMap = sprig.TxtFuncMap()

	for name, fn := range internalFuncMap {
		funcMap[name] = fn
	}
}

// internalFuncMap contains custom functions passed to templates.
// nolint: gochecknoglobals
var internalFuncMap = template.FuncMap{
	"goarch":          goarch,
	"protobuf_goarch": protobufGoarch,
	"protobuf_goos":   protobufGoos,
}

// goarch matches arch representations.
func goarch(s string) string {
	switch s {
	case "386":
		return "i386"

	case "amd64":
		return "x86_64"
	}

	return s
}

// protobuf_goarch matches arch representations.
func protobufGoarch(s string) string {
	switch s {
	case "386":
		return "x86_32"

	case "amd64":
		return "x86_64"
	}

	return s
}

// protobuf_goos matches os representations.
func protobufGoos(s string) string {
	switch s { // nolint: gocritic
	case "darwin":
		return "osx"
	}

	return s
}
