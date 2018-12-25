package gbl

/*
	CREDIT:
	https://blog.gopheracademy.com/advent-2014/parsers-lexers/
*/

import (
	"fmt"
	"strings"

	"github.com/alexanderskafte/go-behave/core"
)

// Parser ...
type Parser struct {
	fnRegistry *Registry
	scanner    *Scanner
	buf        struct {
		tok Token
		lit string
		n   int
	}
	level int
	err   error
}

// NewParser returns a new instance of Parser.
func NewParser(reg *Registry) *Parser {
	return &Parser{fnRegistry: reg}
}

// Compile ...
func (p *Parser) Compile(definition string) (core.Node, error) {
	r := strings.NewReader(definition)
	p.scanner = NewScanner(r)
	return p.parseExpr()
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
	if tok == tokenWhitespace {
		tok, lit = p.scan()
	}
	// fmt.Println(strings.Repeat("   ", p.level) + lit)
	return
}

type item struct {
	tok Token
	lit string
}

var itemInvalid = item{tok: tokenInvalid, lit: "invalid token"}

func (p *Parser) accept(token Token) (item, error) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != token {
		return itemInvalid, expectError(lit, string(token))
	}
	return item{tok, lit}, nil
}

// expectError ...
func expectError(got, exp string) error {
	return fmt.Errorf("got %q, expected %s", got, exp)
}
