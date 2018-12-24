package lang

import (
	"github.com/alexanderskafte/behaviortree/core"
)

func (p *Parser) parseAction(name string) (core.INode, error) {

	if tok, lit := p.scanIgnoreWhitespace(); tok != tokenParenLeft {
		return nil, Error(lit, "(")
	}

	in, err := p.parseList(tokenColon)
	if err != nil {
		return nil, err
	}

	out, err := p.parseList(tokenParenRight)
	if err != nil {
		return nil, err
	}

	tmp, err := p.fnRegistry.GetFunction(name)
	if err != nil {
		return nil, err
	}
	fn := tmp.(core.ActionFn)
	return fn(in, out), nil
}
