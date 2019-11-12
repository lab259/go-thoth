// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: templates/rules/required.gohtml

package rules

import (
	"github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/errors"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
)

// Required generates templates/rules/required.gohtml
func Required(input *RequiredInput) string {
	var _b strings.Builder
	RenderRequired(&_b, input)
	return _b.String()
}

// RenderRequired render templates/rules/required.gohtml
func RenderRequired(_buffer io.StringWriter, input *RequiredInput) {
	switch input.Field.RefType.Name() {
	case "string":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType, *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, input.Field, input.Tag, input.Ref)
		case *myasthurts.StarRefType:
			errors.RenderNil(_buffer, input.Field, input.Tag, input.Ref)
		}
	case "uint":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsUint", input.Field, input.Tag, input.Ref)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, input.Field, input.Tag, input.Ref)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, input.Field, input.Tag, input.Ref)
		}
	case "uint8":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsUint8", input.Field, input.Tag, input.Ref)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, input.Field, input.Tag, input.Ref)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, input.Field, input.Tag, input.Ref)
		}
	case "uint16":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsUint16", input.Field, input.Tag, input.Ref)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, input.Field, input.Tag, input.Ref)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, input.Field, input.Tag, input.Ref)
		}
	case "uint32":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsUint32", input.Field, input.Tag, input.Ref)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, input.Field, input.Tag, input.Ref)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, input.Field, input.Tag, input.Ref)
		}
	case "uint64":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsUint64", input.Field, input.Tag, input.Ref)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, input.Field, input.Tag, input.Ref)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, input.Field, input.Tag, input.Ref)
		}
	case "uintptr":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsUintptr", input.Field, input.Tag, input.Ref)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, input.Field, input.Tag, input.Ref)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, input.Field, input.Tag, input.Ref)
		}
	case "int":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsInt", input.Field, input.Tag, input.Ref)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, input.Field, input.Tag, input.Ref)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, input.Field, input.Tag, input.Ref)
		}
	case "int8":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsInt8", input.Field, input.Tag, input.Ref)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, input.Field, input.Tag, input.Ref)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, input.Field, input.Tag, input.Ref)
		}
	case "int16":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsInt16", input.Field, input.Tag, input.Ref)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, input.Field, input.Tag, input.Ref)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, input.Field, input.Tag, input.Ref)
		}
	case "int32":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsInt32", input.Field, input.Tag, input.Ref)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, input.Field, input.Tag, input.Ref)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, input.Field, input.Tag, input.Ref)
		}
	case "int64":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsInt64", input.Field, input.Tag, input.Ref)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, input.Field, input.Tag, input.Ref)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, input.Field, input.Tag, input.Ref)
		}
	case "float32":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsFloat32", input.Field, input.Tag, input.Ref)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, input.Field, input.Tag, input.Ref)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, input.Field, input.Tag, input.Ref)
		}
	case "float64":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsFloat64", input.Field, input.Tag, input.Ref)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, input.Field, input.Tag, input.Ref)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, input.Field, input.Tag, input.Ref)
		}
	case "complex64":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsComplex64", input.Field, input.Tag, input.Ref)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, input.Field, input.Tag, input.Ref)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, input.Field, input.Tag, input.Ref)
		}
	case "complex128":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsComplex128", input.Field, input.Tag, input.Ref)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, input.Field, input.Tag, input.Ref)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, input.Field, input.Tag, input.Ref)
		}
	case "bool":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType:

			_buffer.WriteString("if ! ")
			_buffer.WriteString(gorazor.HTMLEscape(input.Ref))
			_buffer.WriteString(" {")

			_buffer.WriteString("\terrs = append(errs, NewError(\"")
			_buffer.WriteString(gorazor.HTMLEscape(input.Field.Name))
			_buffer.WriteString("\", \"")
			_buffer.WriteString(gorazor.HTMLEscape(input.Tag.Value))
			_buffer.WriteString("\"))")

			_buffer.WriteString("}")

			if c, ok := Condition[input.Ref]; ok {
				Condition[input.Ref] = (c + " || ! " + input.Ref)
			} else {
				Condition[input.Ref] = (input.Ref + "!= false")
			}
		case *myasthurts.StarRefType:

			_buffer.WriteString("if ")
			_buffer.WriteString(gorazor.HTMLEscape(input.Ref))
			_buffer.WriteString(" == nil || * ")
			_buffer.WriteString(gorazor.HTMLEscape(input.Ref))
			_buffer.WriteString(" == 0{")

			_buffer.WriteString("\terrs = append(errs, NewError(\"")
			_buffer.WriteString(gorazor.HTMLEscape(input.Field.Name))
			_buffer.WriteString("\", \"")
			_buffer.WriteString(gorazor.HTMLEscape(input.Tag.Value))
			_buffer.WriteString("\"))")

			_buffer.WriteString("}")

		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, input.Field, input.Tag, input.Ref)
		}
	default:
		if strings.HasPrefix(input.Field.RefType.Name(), "map") {
			switch input.Field.RefType.(type) {
			case *myasthurts.BaseRefType:
				errors.RenderEmpty(_buffer, input.Field, input.Tag, input.Ref)
			case *myasthurts.StarRefType:
				errors.RenderNil(_buffer, input.Field, input.Tag, input.Ref)
			}
		} else {
			switch input.Field.RefType.(type) {
			case *myasthurts.BaseRefType:
				if input.Field.RefType.Type() != nil {

					_buffer.WriteString("if (")
					_buffer.WriteString(gorazor.HTMLEscape(input.Ref))
					_buffer.WriteString(" == ")
					_buffer.WriteString(gorazor.HTMLEscape(input.Field.RefType.Type().Name()))
					_buffer.WriteString(("{}"))
					_buffer.WriteString(") {")

					_buffer.WriteString("\terrs = append(errs, NewError(\"")
					_buffer.WriteString(gorazor.HTMLEscape(input.Field.Name))
					_buffer.WriteString("\", \"")
					_buffer.WriteString(gorazor.HTMLEscape(input.Tag.Value))
					_buffer.WriteString("\"))")

					_buffer.WriteString("}")

				} else {

					_buffer.WriteString("if IsValid(")
					_buffer.WriteString(gorazor.HTMLEscape(input.Ref))
					_buffer.WriteString(") {")

					_buffer.WriteString("\terrs = append(errs, NewError(\"")
					_buffer.WriteString(gorazor.HTMLEscape(input.Field.Name))
					_buffer.WriteString("\", \"")
					_buffer.WriteString(gorazor.HTMLEscape(input.Tag.Value))
					_buffer.WriteString("\"))")

					_buffer.WriteString("}")

				}
			case *myasthurts.StarRefType:
				errors.RenderNil(_buffer, input.Field, input.Tag, input.Ref)
			case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
				errors.RenderEmpty(_buffer, input.Field, input.Tag, input.Ref)
			}
		}
	}

}
