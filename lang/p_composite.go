package lang

import (
	"github.com/alexanderskafte/behaviortree/core"
)

func (p *Parser) parseComposite(name string) (core.INode, error) {

	if tok, lit := p.scanIgnoreWhitespace(); tok != tokenBracketLeft {
		return nil, Error(lit, "{")
	}

	children, err := p.parseCompositeChildren(tokenBracketRight)
	if err != nil {
		return nil, err
	}

	tmp, err := p.fnRegistry.GetFunction(name)
	if err != nil {
		return nil, err
	}
	fn := tmp.(core.CompositeFn)
	return fn(children...), nil
}

func (p *Parser) parseCompositeChildren(brk Token) ([]core.INode, error) {
	children := []core.INode{}
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
