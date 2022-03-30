package alt

import (
	"errors"
	"fmt"
	"io"
)

// ParserError represents an error that was created by the parser specifically
type ParserError struct {
	msg   string
	where Loc
}

func (e ParserError) Error() string {
	return fmt.Sprintf("%s:%d:%d - %s",
		e.where.File, e.where.Line, e.where.Col, e.msg)
}

// Loc represents a location in an alt file
type Loc struct {
	File string
	Line uint
	Col  uint
}

// TextMode represents a text node's formatting mode
type TextMode uint8

const (
	// TextNormal is no additional formatting
	TextNormal TextMode = 0
	// TextItalic is italicised text, <i>
	TextItalic TextMode = 1
	// TextBold is boldened text, <b>
	TextBold TextMode = 2
	// TextUnder is underlined text, <u>
	TextUnder TextMode = 4
	// TextStrike is struckthrough text, <s>
	TextStrike TextMode = 8
	// TextMark is highlighted text, <mark>
	TextMark TextMode = 16
)

// Node represents a bit of formatted text
type Node struct {
	Mode     TextMode
	Text     string
	Where    Loc
	Disables TextMode
}

// Parser holds the current state of an alt parser
type Parser struct {
	src  io.RuneScanner
	pos  Loc
	last *Node
}

// NewParser creates a new alt parser. The fn argument is used as a filename
// for error reporting.
func NewParser(fn string, src io.RuneScanner) *Parser {
	return &Parser{
		src: src,
		pos: Loc{
			File: fn,
			Line: 1,
			Col:  1,
		},
	}
}

// Next reads and returns the next node in the parser's underlying source.
func (p *Parser) Next() (*Node, error) {
	c, _, err := p.src.ReadRune()
	if err != nil {
		return nil, err
	}
	if c == '(' {
		node, err := p.nextTyped(TextNormal)
		if err != nil {
			return nil, err
		}
		if p.last != nil {
			node.Mode |= p.last.Mode
			node.Mode &^= p.last.Disables
		}
		p.last = node
		return node, err
	}
	err = p.src.UnreadRune()
	if err != nil {
		return nil, err
	}
	node, err := p.nextNormal()
	p.last = node
	return node, err
}

func (p *Parser) nextTyped(mode TextMode) (*Node, error) {
	startLoc := Loc{
		File: p.pos.File,
		Line: p.pos.Line,
		Col:  p.pos.Col,
	}
	c, _, err := p.src.ReadRune()
	if err != nil {
		return nil, err
	}
	nodeMode := mode
	switch c {
	case '(', '/':
		nodeMode |= TextItalic
	case '*':
		nodeMode |= TextBold
	case '_':
		nodeMode |= TextUnder
	case '-':
		nodeMode |= TextStrike
	case '|', '!':
		nodeMode |= TextMark
	default:
		return p.nextFakeTyped(c)
	}
	nodeText := ""
	for {
		c, _, err = p.src.ReadRune()
		if errors.Is(err, io.EOF) {
			return nil, ParserError{
				msg:   fmt.Sprintf("unterminated %d", nodeMode),
				where: startLoc,
			}
		} else if err != nil {
			return nil, err
		}
		if c == '(' {
			err = p.src.UnreadRune()
			if err != nil {
				return nil, err
			}
			return &Node{
				Mode:  nodeMode,
				Text:  nodeText,
				Where: startLoc,
			}, nil
		}
		if c == ')' {
			closing := nodeText[len(nodeText)-1]
			nodeText = nodeText[:len(nodeText)-1]
			disables := TextNormal
			switch closing {
			case ')', '/':
				disables = TextItalic
			case '*':
				disables = TextBold
			case '_':
				disables = TextUnder
			case '-':
				disables = TextStrike
			case '|', '!':
				disables = TextMark
			default:
				nodeText += string(closing)
			}
			if disables != TextNormal {
				return &Node{
					Mode:     nodeMode,
					Text:     nodeText,
					Where:    startLoc,
					Disables: disables,
				}, nil
			}
		}
		nodeText += string(c)
	}
}

func (p *Parser) nextNormal() (*Node, error) {
	startLoc := Loc{
		File: p.pos.File,
		Line: p.pos.Line,
		Col:  p.pos.Col,
	}
	nodeText := ""
	for {
		c, _, err := p.src.ReadRune()
		if errors.Is(err, io.EOF) {
			return &Node{
				Mode:  TextNormal,
				Text:  nodeText,
				Where: startLoc,
			}, nil
		} else if err != nil {
			return nil, err
		}
		if c == '(' {
			err = p.src.UnreadRune()
			if err != nil {
				return nil, err
			}
			return &Node{
				Mode:  TextNormal,
				Text:  nodeText,
				Where: startLoc,
			}, nil
		}
		nodeText += string(c)
	}
}

func (p *Parser) nextFakeTyped(mode rune) (*Node, error) {
	node, err := p.nextNormal()
	if err != nil {
		return nil, err
	}
	node.Text = "(" + string(mode) + node.Text
	node.Where.Col -= 2
	return node, err
}
