package gbl

import (
	"github.com/alexanderskafte/go-behave/core"
)

func (p *Parser) parseDecorator(name string) (core.Node, error) {

	if _, err := p.accept(tokenParenLeft); err != nil {
		return nil, err
	}

	params, err := p.parseAssignmentList(tokenParenRight)
	if err != nil {
		return nil, err
	}

	if _, err := p.accept(tokenBracketLeft); err != nil {
		return nil, err
	}

	child, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	if _, err := p.accept(tokenBracketRight); err != nil {
		return nil, err
	}

	tmp, _, err := p.fnRegistry.Get(name)
	if err != nil {
		return nil, err
	}
	fn := tmp.(core.DecoratorFn)
	return fn(params, child), nil
}
