package alt

import (
	"bytes"
	"io"
	"strings"
)

// RenderFile renders alt-formatted text to HTML from the given RuneScanner.
// The fn argument is a filename to use for error messages.
func RenderFile(src io.RuneScanner, fn string) ([]byte, error) {
	p := NewParser(fn, src)
	html := bytes.NewBuffer(make([]byte, 0))
	_, err := GenerateHTML(p, html)
	if err != nil {
		return nil, err
	}
	return html.Bytes(), nil
}

// RenderString renders alt-formatted text to HTML form the given string.
// The name is used for error reporting.
func RenderString(src, name string) ([]byte, error) {
	rdr := strings.NewReader(src)
	p := NewParser(name, rdr)
	html := bytes.NewBuffer(make([]byte, 0))
	_, err := GenerateHTML(p, html)
	if err != nil {
		return nil, err
	}
	return html.Bytes(), nil
}
