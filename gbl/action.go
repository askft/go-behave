package gbl

import (
	"github.com/alexanderskafte/go-behave/core"
)

func (p *Parser) parseAction(name string) (core.Node, error) {

	if tok, lit := p.scanIgnoreWhitespace(); tok != tokenParenLeft {
		return nil, expectError(lit, "(")
	}

	params, err := p.parseAssignmentList(tokenColon)
	if err != nil {
		return nil, err
	}

	returns, err := p.parseAssignmentList(tokenParenRight)
	if err != nil {
		return nil, err
	}

	tmp, _, err := p.fnRegistry.Get(name)
	if err != nil {
		return nil, err
	}
	fn := tmp.(core.ActionFn)
	return fn(params, returns), nil
}
