// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: templates/errors/empty.gohtml

package errors

import (
	"github.com/lab259/go-my-ast-hurts"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
)

// Empty generates templates/errors/empty.gohtml
func Empty(field *myasthurts.Field, tag myasthurts.TagParam, value interface{}) string {
	var _b strings.Builder
	RenderEmpty(&_b, field, tag, value)
	return _b.String()
}

// RenderEmpty render templates/errors/empty.gohtml
func RenderEmpty(_buffer io.StringWriter, field *myasthurts.Field, tag myasthurts.TagParam, value interface{}) {
	_buffer.WriteString("\nif Empty(len(")
	_buffer.WriteString(gorazor.HTMLEscape(value))
	_buffer.WriteString(")) {\n\terrs = append(errs, ErrEmpty(\"")
	_buffer.WriteString(gorazor.HTMLEscape(field.Name))
	_buffer.WriteString("\", \"")
	_buffer.WriteString(gorazor.HTMLEscape(tag.Value))
	_buffer.WriteString("\"))\n}")

}
