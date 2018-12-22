package lang

import (
	"github.com/alexanderskafte/behaviortree/core"
)

func (p *Parser) parseComposite(name string) (core.INode, error) {

	if tok, lit := p.scanIgnoreWhitespace(); tok != tokenBL {
		return nil, Error(lit, "{")
	}

	children, err := p.parseCompositeChildren(tokenBR)
	if err != nil {
		return nil, err
	}

	base := core.NewComposite()
	base.AddChildren(children...)

	return p.nodeRegistry.NewComposite(name, base)
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
