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
func Required(field *myasthurts.Field, tag myasthurts.TagParam, value interface{}) string {
	var _b strings.Builder
	RenderRequired(&_b, field, tag, value)
	return _b.String()
}

// RenderRequired render templates/rules/required.gohtml
func RenderRequired(_buffer io.StringWriter, field *myasthurts.Field, tag myasthurts.TagParam, value interface{}) {
	switch field.RefType.Name() {
	case "string":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType, *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, field, tag, value)
		case *myasthurts.StarRefType:
			errors.RenderNil(_buffer, field, tag, value)
		}
	case "uint":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsUint", field, tag, value)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, field, tag, value)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, field, tag, value)
		}
	case "uint8":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsUint8", field, tag, value)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, field, tag, value)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, field, tag, value)
		}
	case "uint16":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsUint16", field, tag, value)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, field, tag, value)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, field, tag, value)
		}
	case "uint32":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsUint32", field, tag, value)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, field, tag, value)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, field, tag, value)
		}
	case "uint64":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsUint64", field, tag, value)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, field, tag, value)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, field, tag, value)
		}
	case "uintptr":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsUintptr", field, tag, value)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, field, tag, value)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, field, tag, value)
		}
	case "int":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsInt", field, tag, value)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, field, tag, value)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, field, tag, value)
		}
	case "int8":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsInt8", field, tag, value)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, field, tag, value)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, field, tag, value)
		}
	case "int16":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsInt16", field, tag, value)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, field, tag, value)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, field, tag, value)
		}
	case "int32":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsInt32", field, tag, value)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, field, tag, value)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, field, tag, value)
		}
	case "int64":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsInt64", field, tag, value)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, field, tag, value)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, field, tag, value)
		}
	case "float32":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsFloat32", field, tag, value)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, field, tag, value)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, field, tag, value)
		}
	case "float64":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsFloat64", field, tag, value)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, field, tag, value)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, field, tag, value)
		}
	case "complex64":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsComplex64", field, tag, value)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, field, tag, value)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, field, tag, value)
		}
	case "complex128":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType:
			errors.RenderNumber(_buffer, "IsComplex128", field, tag, value)
		case *myasthurts.StarRefType:
			errors.RenderZero(_buffer, field, tag, value)
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			errors.RenderEmpty(_buffer, field, tag, value)
		}
	default:
		if strings.HasPrefix(field.RefType.Name(), "map") {
			switch field.RefType.(type) {
			case *myasthurts.BaseRefType:
				errors.RenderEmpty(_buffer, field, tag, value)
			case *myasthurts.StarRefType:
				errors.RenderNil(_buffer, field, tag, value)
			}
		} else {
			switch field.RefType.(type) {
			case *myasthurts.BaseRefType:
				if field.RefType.Type() != nil {

					_buffer.WriteString("if (")
					_buffer.WriteString(gorazor.HTMLEscape(value))
					_buffer.WriteString(" == ")
					_buffer.WriteString(gorazor.HTMLEscape(field.RefType.Type().Name()))
					_buffer.WriteString(("{}"))
					_buffer.WriteString(") {")

					_buffer.WriteString("\terrs = append(errs, NewError(\"")
					_buffer.WriteString(gorazor.HTMLEscape(field.Name))
					_buffer.WriteString("\", \"")
					_buffer.WriteString(gorazor.HTMLEscape(tag.Value))
					_buffer.WriteString("\"))")

					_buffer.WriteString("}")

				} else {

					_buffer.WriteString("if IsValid(")
					_buffer.WriteString(gorazor.HTMLEscape(value))
					_buffer.WriteString(") {")

					_buffer.WriteString("\terrs = append(errs, NewError(\"")
					_buffer.WriteString(gorazor.HTMLEscape(field.Name))
					_buffer.WriteString("\", \"")
					_buffer.WriteString(gorazor.HTMLEscape(tag.Value))
					_buffer.WriteString("\"))")

					_buffer.WriteString("}")

				}
			case *myasthurts.StarRefType:
				errors.RenderNil(_buffer, field, tag, value)
			case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
				errors.RenderEmpty(_buffer, field, tag, value)
			}
		}
	}

}
