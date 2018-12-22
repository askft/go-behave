package lang

/*
	CREDIT:
	https://blog.gopheracademy.com/advent-2014/parsers-lexers/
*/

import (
	"fmt"
	"io"
	"strings"

	"github.com/alexanderskafte/behaviortree/core"
	"github.com/alexanderskafte/behaviortree/registry"
)

// Parser ...
type Parser struct {
	nodeRegistry registry.Registry
	scanner      *Scanner
	buf          struct {
		tok Token
		lit string
		n   int
	}
	level int
	err   error
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader, reg registry.Registry) *Parser {
	return &Parser{scanner: NewScanner(r), nodeRegistry: reg}
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
	tok, lit = p.scanner.Scan()
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
	if tok == tokenWS {
		tok, lit = p.scan()
	}
	fmt.Println(strings.Repeat("   ", p.level) + lit)
	return
}

// --------------------------------------------------------
// BehaviorTree parsing functions

// Parse ...
func (p *Parser) Parse() (core.INode, error) {
	return p.parseExpr()
}

func (p *Parser) parseExpr() (core.INode, error) {
	p.level++
	defer func() { p.level-- }()

	tok, lit := p.scanIgnoreWhitespace()
	if !isKeyword(lit) {
		return nil, Error(lit, "keyword")
	}

	idTok, name := p.scanIgnoreWhitespace()
	if idTok != tokenID {
		return nil, Error(name, "identifier")
	}

	switch tok {
	case tokenComposite:
		return p.parseComposite(name)
	case tokenDecorator:
		return p.parseDecorator(name)
	case tokenCondition:
		return p.parseCondition(name)
	case tokenAction:
		return p.parseAction(name)
	}
	return nil, fmt.Errorf("invalid token")
}

// Error ...
func Error(got, exp string) error {
	return fmt.Errorf("got %q, expected %q", got, exp)
}
