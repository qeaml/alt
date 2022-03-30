package alt_test

import (
	"errors"
	"io"
	"strings"
	"testing"

	. "github.com/qeaml/alt"
)

func TestNormal(t *testing.T) {
	src := strings.NewReader("hello")
	p := NewParser("TestNormal", src)
	node, err := p.Next()
	if err != nil {
		t.Fatal(err)
	}
	if node.Mode != 0 {
		t.Fatalf("node mode should be normal, but isn't: %d", node.Mode)
	}
	if len(node.Text) != 5 {
		t.Fatalf("node text is incorrect: %s", node.Text)
	}
	node, err = p.Next()
	if !errors.Is(err, io.EOF) {
		t.Fatalf("expected end of file, but got extra node: %d `%s`", node.Mode, node.Text)
	}
}

func TestItalic(t *testing.T) {
	src := strings.NewReader("((hello))")
	p := NewParser("TestItalic", src)
	node, err := p.Next()
	if err != nil {
		t.Fatal(err)
	}
	if node.Mode&TextItalic == 0 {
		t.Fatalf("node mode should be italic, but isn't: %d", node.Mode)
	}
	if len(node.Text) != 5 {
		t.Fatalf("node text is incorrect: %s", node.Text)
	}
	node, err = p.Next()
	if !errors.Is(err, io.EOF) {
		t.Fatalf("expected end of file, but got extra node: %d `%s`", node.Mode, node.Text)
	}
}

func TestBold(t *testing.T) {
	src := strings.NewReader("(*hello*)")
	p := NewParser("TestBold", src)
	node, err := p.Next()
	if err != nil {
		t.Fatal(err)
	}
	if node.Mode&TextBold == 0 {
		t.Fatalf("node mode should be bold, but isn't: %d", node.Mode)
	}
	if len(node.Text) != 5 {
		t.Fatalf("node text is incorrect: %s", node.Text)
	}
	node, err = p.Next()
	if !errors.Is(err, io.EOF) {
		t.Fatalf("expected end of file, but got extra node: %d `%s`", node.Mode, node.Text)
	}
}

func TestFake(t *testing.T) {
	src := strings.NewReader("(yeah)")
	p := NewParser("TestFake", src)
	node, err := p.Next()
	if err != nil {
		t.Fatal(err)
	}
	if node.Mode != TextNormal {
		t.Fatalf("node mode should be normal, but isn't: %d", node.Mode)
	}
	if len(node.Text) != len("(yeah)") {
		t.Fatalf("node text is incorrect: %s", node.Text)
	}
	node, err = p.Next()
	if !errors.Is(err, io.EOF) {
		t.Fatalf("expected end of file, but got extra node: %d `%s`", node.Mode, node.Text)
	}
}

func TestMulti(t *testing.T) {
	src := strings.NewReader("(*hello*)((world))")
	p := NewParser("TestMulti", src)
	node, err := p.Next()
	if err != nil {
		t.Fatal(err)
	}
	if node.Mode&TextBold == 0 {
		t.Fatalf("node 0 mode should be bold, but isn't: %d", node.Mode)
	}
	if len(node.Text) != 5 {
		t.Fatalf("node 0 text is incorrect: %s", node.Text)
	}
	node, err = p.Next()
	if err != nil {
		t.Fatal(err)
	}
	if node.Mode&TextItalic == 0 {
		t.Fatalf("node 1 mode should be italic, but isn't: %d", node.Mode)
	}
	if len(node.Text) != 5 {
		t.Fatalf("node 1 text is incorrect: %s", node.Text)
	}
	node, err = p.Next()
	if !errors.Is(err, io.EOF) {
		t.Fatalf("expected end of file, but got extra node: %d `%s`", node.Mode, node.Text)
	}
}

func TestMixed(t *testing.T) {
	src := strings.NewReader("(*hello*) there ((world))")
	p := NewParser("TestMixed", src)
	node, err := p.Next()
	if err != nil {
		t.Fatal(err)
	}
	if node.Mode&TextBold == 0 {
		t.Fatalf("node 0 mode should be bold, but isn't: %d", node.Mode)
	}
	if len(node.Text) != 5 {
		t.Fatalf("node 0 text is incorrect: %s", node.Text)
	}
	node, err = p.Next()
	if err != nil {
		t.Fatal(err)
	}
	if node.Mode != TextNormal {
		t.Fatalf("node 1 mode should be normal, but isn't: %d", node.Mode)
	}
	if len(node.Text) != len(" there ") {
		t.Fatalf("node 1 text incorrect: `%s`", node.Text)
	}
	node, err = p.Next()
	if err != nil {
		t.Fatal(err)
	}
	if node.Mode&TextItalic == 0 {
		t.Fatalf("node 2 mode should be italic, but isn't: %d", node.Mode)
	}
	if len(node.Text) != 5 {
		t.Fatalf("node 2 text is incorrect: %s", node.Text)
	}
	node, err = p.Next()
	if !errors.Is(err, io.EOF) {
		t.Fatalf("expected end of file, but got extra node: %d `%s`", node.Mode, node.Text)
	}
}

func TestMixedFake(t *testing.T) {
	src := strings.NewReader("(*hello*) (there) ((world))")
	p := NewParser("TestMixedFake", src)
	node, err := p.Next()
	if err != nil {
		t.Fatal(err)
	}
	if node.Mode&TextBold == 0 {
		t.Fatalf("node 0 mode should be bold, but isn't: %d", node.Mode)
	}
	if len(node.Text) != 5 {
		t.Fatalf("node 0 text is incorrect: %s", node.Text)
	}
	node, err = p.Next()
	if err != nil {
		t.Fatal(err)
	}
	if node.Mode != TextNormal {
		t.Fatalf("node 1 mode should be normal, but isn't: %d", node.Mode)
	}
	if len(node.Text) != len(" ") {
		t.Fatalf("node 1 text incorrect: `%s`", node.Text)
	}
	node, err = p.Next()
	if err != nil {
		t.Fatal(err)
	}
	if node.Mode != TextNormal {
		t.Fatalf("node 2 mode should be normal, but isn't: %d", node.Mode)
	}
	if len(node.Text) != len("(there) ") {
		t.Fatalf("node 2 text incorrect: `%s`", node.Text)
	}
	node, err = p.Next()
	if err != nil {
		t.Fatal(err)
	}
	if node.Mode&TextItalic == 0 {
		t.Fatalf("node 3 mode should be italic, but isn't: %d", node.Mode)
	}
	if len(node.Text) != 5 {
		t.Fatalf("node 3 text is incorrect: %s", node.Text)
	}
	node, err = p.Next()
	if !errors.Is(err, io.EOF) {
		t.Fatalf("expected end of file, but got extra node: %d `%s`", node.Mode, node.Text)
	}
}

func TestNested(t *testing.T) {
	src := strings.NewReader("(*hello ((there))*)")
	p := NewParser("TestNested", src)
	node, err := p.Next()
	if err != nil {
		t.Fatal(err)
	}
	if node.Mode&TextBold == 0 {
		t.Fatalf("node 0 mode should be bold, but isn't: %d", node.Mode)
	}
	if len(node.Text) != len("hello ") {
		t.Fatalf("node 0 text is incorrect: %s", node.Text)
	}
	node, err = p.Next()
	if err != nil {
		t.Fatal(err)
	}
	if node.Mode&TextBold == 0 {
		t.Fatalf("node 1 mode should be bold, but isn't: %d", node.Mode)
	}
	if node.Mode&TextItalic == 0 {
		t.Fatalf("node 1 mode should be italic, but isn't: %d", node.Mode)
	}
	if len(node.Text) != len("there") {
		t.Fatalf("node 0 text is incorrect: %s", node.Text)
	}
	node, err = p.Next()
	if !errors.Is(err, io.EOF) {
		t.Fatalf("expected end of file, but got extra node: %d `%s` (disables %d)",
			node.Mode, node.Text, node.Disables)
	}
}
