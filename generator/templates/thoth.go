// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: templates/thoth.gohtml

package templates

import (
	"github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
)

// Thoth generates templates/thoth.gohtml
func Thoth(fileName string, pkg *myasthurts.Package, structsThoth []*myasthurts.Struct) string {
	var _b strings.Builder
	RenderThoth(&_b, fileName, pkg, structsThoth)
	return _b.String()
}

// RenderThoth render templates/thoth.gohtml
func RenderThoth(_buffer io.StringWriter, fileName string, pkg *myasthurts.Package, structsThoth []*myasthurts.Struct) {
	_buffer.WriteString("\npackage ")
	_buffer.WriteString(gorazor.HTMLEscape(pkg.Name))
	if hasTag(structsThoth) {
		for _, s := range structsThoth {
			// Clean conditions
			rules.MapCondition = make(map[string]string, 10)

			_buffer.WriteString(("\n// Validate TODO\n"))

			strRef := strings.ToLower(s.Name()[0:1])

			_buffer.WriteString("func(")
			_buffer.WriteString(gorazor.HTMLEscape(strRef))
			_buffer.WriteString(" ")
			_buffer.WriteString(("*"))
			_buffer.WriteString(gorazor.HTMLEscape(s.Name()))
			_buffer.WriteString(") Validate() (errs ValidationErrors) {")

			for _, field := range s.Fields {
				for _, tag := range field.Tag.Params {
					attribute := strRef + "." + field.Name
					filterValidate(_buffer, &FilterInput{
						Struct:    s,
						StructRef: strRef,
						Field:     field,
						Tag:       tag,
						Ref:       attribute,
					})

					_buffer.WriteString(("\n"))

				} // For tags
			} // For fields

			_buffer.WriteString("return errs")

			_buffer.WriteString("}")

		}

	}

}
