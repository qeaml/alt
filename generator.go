package alt

import (
	"errors"
	"io"
)

// GenerateHTML reads text nodes from the given parser, and formats them
// into the corresponding HTML to the out writer.
func GenerateHTML(src *Parser, out io.Writer) (n int, err error) {
	for {
		node, err := src.Next()
		if errors.Is(err, io.EOF) {
			return n, nil
		} else if err != nil {
			return n, err
		}
		if node.Mode&TextItalic > 0 {
			amt, err := out.Write([]byte("<i>"))
			n += amt
			if err != nil {
				return n, err
			}
		}
		if node.Mode&TextBold > 0 {
			amt, err := out.Write([]byte("<b>"))
			n += amt
			if err != nil {
				return n, err
			}
		}
		if node.Mode&TextUnder > 0 {
			amt, err := out.Write([]byte("<u>"))
			n += amt
			if err != nil {
				return n, err
			}
		}
		if node.Mode&TextStrike > 0 {
			amt, err := out.Write([]byte("<s>"))
			n += amt
			if err != nil {
				return n, err
			}
		}
		if node.Mode&TextMark > 0 {
			amt, err := out.Write([]byte("<mark>"))
			n += amt
			if err != nil {
				return n, err
			}
		}
		amt, err := out.Write([]byte(node.Text))
		n += amt
		if err != nil {
			return n, err
		}
		if node.Mode&TextItalic > 0 {
			amt, err := out.Write([]byte("</i>"))
			n += amt
			if err != nil {
				return n, err
			}
		}
		if node.Mode&TextBold > 0 {
			amt, err := out.Write([]byte("</b>"))
			n += amt
			if err != nil {
				return n, err
			}
		}
		if node.Mode&TextUnder > 0 {
			amt, err := out.Write([]byte("</u>"))
			n += amt
			if err != nil {
				return n, err
			}
		}
		if node.Mode&TextStrike > 0 {
			amt, err := out.Write([]byte("</s>"))
			n += amt
			if err != nil {
				return n, err
			}
		}
		if node.Mode&TextMark > 0 {
			amt, err := out.Write([]byte("</mark>"))
			n += amt
			if err != nil {
				return n, err
			}
		}
	}
}

// GenerateHTML reads text nodes from the given parser, and formats them
// into the corresponding Markdown to the out writer.
func GenerateMarkdown(src *Parser, out io.Writer) (n int, err error) {
	for {
		node, err := src.Next()
		if errors.Is(err, io.EOF) {
			return n, nil
		} else if err != nil {
			return n, err
		}
		if node.Mode&TextItalic > 0 {
			amt, err := out.Write([]byte("*"))
			n += amt
			if err != nil {
				return n, err
			}
		}
		if node.Mode&TextBold > 0 {
			amt, err := out.Write([]byte("**"))
			n += amt
			if err != nil {
				return n, err
			}
		}
		if node.Mode&TextUnder > 0 {
			amt, err := out.Write([]byte("__"))
			n += amt
			if err != nil {
				return n, err
			}
		}
		if node.Mode&TextStrike > 0 {
			amt, err := out.Write([]byte("~~"))
			n += amt
			if err != nil {
				return n, err
			}
		}
		// note: markdown doesn't have an equivalent for Mark :)
		amt, err := out.Write([]byte(node.Text))
		n += amt
		if err != nil {
			return n, err
		}
		if node.Mode&TextItalic > 0 {
			amt, err := out.Write([]byte("*"))
			n += amt
			if err != nil {
				return n, err
			}
		}
		if node.Mode&TextBold > 0 {
			amt, err := out.Write([]byte("**"))
			n += amt
			if err != nil {
				return n, err
			}
		}
		if node.Mode&TextUnder > 0 {
			amt, err := out.Write([]byte("__"))
			n += amt
			if err != nil {
				return n, err
			}
		}
		if node.Mode&TextStrike > 0 {
			amt, err := out.Write([]byte("~~"))
			n += amt
			if err != nil {
				return n, err
			}
		}
	}
}
