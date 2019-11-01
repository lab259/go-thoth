// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: templates/struct.gohtml

package templates

import (
	"github.com/lab259/go-my-ast-hurts"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
)

// Struct generates templates/struct.gohtml
func Struct(fileName string, pkg *myasthurts.Package, structsThoth []*myasthurts.Struct) string {
	var _b strings.Builder
	RenderStruct(&_b, fileName, pkg, structsThoth)
	return _b.String()
}

// RenderStruct render templates/struct.gohtml
func RenderStruct(_buffer io.StringWriter, fileName string, pkg *myasthurts.Package, structsThoth []*myasthurts.Struct) {
	_buffer.WriteString(("//go:generate gofmt -s -w"))
	_buffer.WriteString(" ")
	_buffer.WriteString(gorazor.HTMLEscape(fileName))
	_buffer.WriteString("\npackage ")
	_buffer.WriteString(gorazor.HTMLEscape(pkg.Name))
	_buffer.WriteString(" ")
	_buffer.WriteString(("\n\n"))
	for _, s := range structsThoth {

		_buffer.WriteString(("\n// Thoth validate\n"))

		_buffer.WriteString("func(")
		_buffer.WriteString(gorazor.HTMLEscape(strings.ToLower(s.Name()[0:1])))
		_buffer.WriteString(" ")
		_buffer.WriteString(("*"))
		_buffer.WriteString(gorazor.HTMLEscape(s.Name()))
		_buffer.WriteString(") Thoth() error {")

		for _, field := range s.Fields {
			for _, tag := range field.Tag.Params {
				if tag.Value == "required" {

					_buffer.WriteString("if required(\"")
					_buffer.WriteString(gorazor.HTMLEscape(field.RefType.Name()))
					_buffer.WriteString("\", ")
					_buffer.WriteString(gorazor.HTMLEscape(strings.ToLower(s.Name()[0:1])))
					_buffer.WriteString(".")
					_buffer.WriteString(gorazor.HTMLEscape(field.Name))
					_buffer.WriteString(") {")

					_buffer.WriteString(("return errRequiredString"))

					_buffer.WriteString("}")

					_buffer.WriteString(("\n"))

				}

			}
		}

		_buffer.WriteString(("return nil"))

		_buffer.WriteString("}")

	}

}
