// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: generator/template/struct.gohtml

package template

import (
	"github.com/lab259/go-my-ast-hurts"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
)

// Struct generates generator/template/struct.gohtml
func Struct(pkg *myasthurts.Package, structsThoth []*myasthurts.Struct) string {
	var _b strings.Builder
	RenderStruct(&_b, pkg, structsThoth)
	return _b.String()
}

// RenderStruct render generator/template/struct.gohtml
func RenderStruct(_buffer io.StringWriter, pkg *myasthurts.Package, structsThoth []*myasthurts.Struct) {
	_buffer.WriteString("\npackage ")
	_buffer.WriteString(gorazor.HTMLEscStr(pkg.Name))
	_buffer.WriteString(" ")
	_buffer.WriteString(("\n\n"))
	for _, s := range structsThoth {

		_buffer.WriteString(("// Thoth description\n"))

		_buffer.WriteString("func(")
		_buffer.WriteString(gorazor.HTMLEscStr(strings.ToLower(s.Name()[0:1])))
		_buffer.WriteString(" ")
		_buffer.WriteString(("*"))
		_buffer.WriteString(gorazor.HTMLEscStr(s.Name()))
		_buffer.WriteString(") Thoth() error { ")
		_buffer.WriteString(("\n return nil \n}\n"))

	}

}
