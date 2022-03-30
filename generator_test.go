package alt_test

import (
	"bytes"
	"strings"
	"testing"

	. "github.com/qeaml/alt"
)

func TestHTML(t *testing.T) {
	src := strings.NewReader("hello ((world))")
	p := NewParser("TestHTML", src)
	out := bytes.NewBuffer(make([]byte, 0))
	n, err := GenerateHTML(p, out)
	if err != nil {
		t.Fatal(err)
	}
	html := string(out.Bytes()[:n])
	if !strings.Contains(html, "<i>") {
		t.Fatalf("resulting HTML has no <i> tag: %s", html)
	}
	if !strings.Contains(html, "</i>") {
		t.Fatalf("resulting HTML has no </i> tag: %s", html)
	}
}

func TestMarkdown(t *testing.T) {
	src := strings.NewReader("hello (*world*)")
	p := NewParser("TestMarkdown", src)
	out := bytes.NewBuffer(make([]byte, 0))
	n, err := GenerateMarkdown(p, out)
	if err != nil {
		t.Fatal(err)
	}
	md := string(out.Bytes()[:n])
	if !strings.Contains(md, "**") {
		t.Fatalf("resulting Markdown has no ** tag: %s", md)
	}
}
