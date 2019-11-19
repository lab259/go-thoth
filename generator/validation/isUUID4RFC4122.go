package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsUUID4RFC4122Input TODO
type IsUUID4RFC4122Input struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsUUID4RFC4122 TODO
func IsUUID4RFC4122(_buffer io.StringWriter, input *IsUUID4RFC4122Input) {
	condition, isLoop := isRegex("uUID4RFC4122Regex", input.Field, input.Ref, "uuid4_rfc4122")
	if isLoop {
		rules.RenderLoop(
			_buffer,
			condition,
			input.Ref,
			input.Field,
			input.Tag,
		)
	} else {
		rules.RenderCondition(
			_buffer,
			condition,
			input.Field,
			input.Tag,
		)
	}
}
