package gbl

import (
	"github.com/alexanderskafte/go-behave/core"
)

func (p *Parser) parseAction(name string) (core.Node, error) {

	if tok, lit := p.scanIgnoreWhitespace(); tok != tokenParenLeft {
		return nil, expectError(lit, "(")
	}

	in, err := p.parseList(tokenColon)
	if err != nil {
		return nil, err
	}

	out, err := p.parseList(tokenParenRight)
	if err != nil {
		return nil, err
	}

	tmp, _, err := p.fnRegistry.Get(name)
	if err != nil {
		return nil, err
	}
	fn := tmp.(core.ActionFn)
	return fn(in, out), nil
}
