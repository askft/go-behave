package gbl

import (
	"github.com/alexanderskafte/go-behave/core"
)

func (p *Parser) parseComposite(name string) (core.Node, error) {

	if tok, lit := p.scanIgnoreWhitespace(); tok != tokenBracketLeft {
		return nil, expectError(lit, "{")
	}

	children, err := p.parseCompositeChildren(tokenBracketRight)
	if err != nil {
		return nil, err
	}

	tmp, _, err := p.fnRegistry.Get(name)
	if err != nil {
		return nil, err
	}
	fn := tmp.(core.CompositeFn)
	return fn(children...), nil
}

func (p *Parser) parseCompositeChildren(brk Token) ([]core.Node, error) {
	children := []core.Node{}
	for {
		tok, _ := p.scanIgnoreWhitespace()
		if tok == brk {
			break
		} else {
			p.unscan()
		}
		child, err := p.parseExpr()
		if err != nil {
			return nil, err
		}
		children = append(children, child)
	}
	return children, nil
}
