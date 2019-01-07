// nolint: goconst
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
	"jq_goarch":       jqGoarch,
	"jq_osarch":       jqOsarch,
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

// jq_goarch matches arch representations.
func jqGoarch(s string) string {
	switch s {
	case "386":
		return "x86"

	case "amd64":
		return "x86_64"
	}

	return s
}

// jq_osarch matches arch representations.
func jqOsarch(os string, arch string) string {
	switch os {
	case "linux":
		switch arch {
		case "386":
			return "linux32"

		case "amd64":
			return "linux64"
		}

	case "darwin":
		return "osx-amd64"

	case "windows":
		switch arch {
		case "386":
			return "win32.exe"

		case "amd64":
			return "win64.exe"
		}
	}

	return ""
}
