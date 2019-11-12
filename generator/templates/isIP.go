// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: templates/isIP.gohtml

package templates

import (
	"io"
	"strings"
)

// IsIP generates templates/isIP.gohtml
func IsIP() string {
	var _b strings.Builder
	RenderIsIP(&_b)
	return _b.String()
}

// RenderIsIP render templates/isIP.gohtml
func RenderIsIP(_buffer io.StringWriter) {
	_buffer.WriteString("\n// IsIP is the validation function for validating if the field's value is a valid v4 or v6 IP address.\nfunc isIP(ref string) bool {\n\tip := net.ParseIP(ref)\n\n\treturn ip != nil\n}\n\n// IsIPv4 is the validation function for validating if a value is a valid v4 IP address.\nfunc isIPv4(ref string) bool {\n    ip := net.ParseIP(ref)\n\n\treturn ip != nil && ip.To4() == nil\n}\n\n// IsIPv6 is the validation function for validating if the field's value is a valid v6 IP address.\nfunc isIPv6(ref string) bool {\n    ip := net.ParseIP(ref)\n\n    // TODO\n\treturn ip != nil && ip.To16() == nil\n}")

}
