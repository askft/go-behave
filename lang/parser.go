package lang

/*
	CREDIT:
	https://blog.gopheracademy.com/advent-2014/parsers-lexers/
*/

import (
	"fmt"
	"io"
	"strings"

	"github.com/alexanderskafte/behaviortree/composite"
	"github.com/alexanderskafte/behaviortree/core"
	"github.com/alexanderskafte/behaviortree/decorator"
)

// Parser ...
type Parser struct {
	s   *Scanner
	buf struct {
		tok Token
		lit string
		n   int
	}
	level int
	err   error
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

// scan returns the next Token from the underlying scanner.
// If a Token has been unscanned then read that instead.
func (p *Parser) scan() (tok Token, lit string) {
	// If we have a Token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next Token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// unscan pushes the previously read Token back onto the buffer.
func (p *Parser) unscan() {
	p.buf.n = 1
	// fmt.Println("[unscan]")
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == tWS {
		tok, lit = p.scan()
	}
	fmt.Println(strings.Repeat("   ", p.level) + lit)
	return
}

// --------------------------------------------------------
// BehaviorTree parsing functions

// Parse ...
func (p *Parser) Parse() (core.INode, error) {
	node, err := p.parse()
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (p *Parser) parse() (core.INode, error) {
	return p.parseExpr()
}

func (p *Parser) parseExpr() (core.INode, error) {
	tok, lit := p.scanIgnoreWhitespace()
	if !isKeyword(lit) {
		return nil, Error(lit, "keyword")
	}
	p.level++
	var (
		node core.INode
		err  error
	)
	switch tok {
	case tSEQUENCE, tSELECTOR:
		node, err = p.parseComposite(tok)
		if err != nil {
			return nil, err
		}
	case tINVERTER:
		node, err = p.parseDecorator(tok)
		if err != nil {
			return nil, err
		}
	case tACTION:
		node, err = p.parseLeaf()
		if err != nil {
			return nil, err
		}
	}
	p.level--
	return node, nil
}

func (p *Parser) parseComposite(comp Token) (core.INode, error) {
	if tok, lit := p.scanIgnoreWhitespace(); tok != tBL {
		return nil, Error(lit, "{")
	}

	node := core.NewComposite()

	for {
		tok, _ := p.scanIgnoreWhitespace()
		if tok == tBR { // end of list
			break
		} else {
			p.unscan()
		}

		child, err := p.parseExpr()
		if err != nil {
			return nil, err
		}
		node.AddChildren(child)
	}

	var specnode core.INode
	switch comp {
	case tSEQUENCE:
		node.Type = core.TypeSequence // TODO remove
		specnode = &composite.Sequence{Composite: node}
	case tSELECTOR:
		node.Type = core.TypeSelector // TODO remove
		specnode = &composite.Selector{node}
	default:
		return nil, fmt.Errorf("invalid composite type")
	}

	return specnode, nil
}

func (p *Parser) parseDecorator(deco Token) (core.INode, error) {

	if tok, lit := p.scanIgnoreWhitespace(); tok != tBL {
		return nil, Error(lit, "{")
	}

	child, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	node := core.NewDecorator()
	node.SetChild(child)

	var specnode core.INode
	switch deco {
	case tINVERTER:
		node.Type = core.TypeInverter
		specnode = &decorator.Inverter{node}
	default:
		return nil, fmt.Errorf("invalid decorator type")
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != tBR {
		return nil, Error(lit, "}")
	}

	return specnode, nil
}

func (p *Parser) parseLeaf() (core.INode, error) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != tID {
		return nil, Error(lit, "identifier")
	}

	node := core.NewLeaf(core.Type(lit), lit)

	return node, nil
}

// Error ...
func Error(got, exp string) error {
	return fmt.Errorf("got %q, expected %s", got, exp)
}

// func (p *Parser) safeScanKeyword() {
// 	if p.err != nil {
// 		return
// 	}
// 	tok, lit := p.scanIgnoreWhitespace()
// 	if !isKeyword(lit) {
// 		p.err = Error(lit, "keyword")
// 	}
// }
