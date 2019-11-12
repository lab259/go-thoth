// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: templates/isFile.gohtml

package templates

import (
	"io"
	"strings"
)

// IsFile generates templates/isFile.gohtml
func IsFile() string {
	var _b strings.Builder
	RenderIsFile(&_b)
	return _b.String()
}

// RenderIsFile render templates/isFile.gohtml
func RenderIsFile(_buffer io.StringWriter) {
	_buffer.WriteString("\nfunc isFile(ref string) bool {\n\tfileInfo, err := os.Stat(ref)\n    if err != nil {\n        return false\n    }\n\n    return !fileInfo.IsDir()\n}\n\n// IsDir is the validation function for validating if the current field's value is a valid directory.\nfunc isDir(dir string) bool {\n    fileInfo, err := os.Stat(dir)\n    if err != nil {\n        return false\n    }\n\n    return fileInfo.IsDir()\n}")

}
