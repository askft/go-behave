package lang

import (
	"github.com/alexanderskafte/behaviortree/core"
)

func (p *Parser) parseDecorator(name string) (core.INode, error) {

	if tok, lit := p.scanIgnoreWhitespace(); tok != tokenBL {
		return nil, Error(lit, "{")
	}

	child, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != tokenBR {
		return nil, Error(lit, "}")
	}

	base := core.NewDecorator()
	base.Child = child

	return p.nodeRegistry.NewDecorator(name, base)
}
