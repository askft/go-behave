package lang

import (
	"github.com/alexanderskafte/behaviortree/core"
)

func (p *Parser) parseCondition(name string) (core.INode, error) {

	panic("not implemented")

	// tok, lit := p.scanIgnoreWhitespace()
	// if tok != tokenID {
	// 	return nil, Error(lit, "identifier")
	// }
	// function := lit

	// tok, lit = p.scanIgnoreWhitespace()
	// if tok != tokenPL {
	// 	return nil, Error(lit, "(")
	// }

	// in := []string{}

	// // Parse the input list
	// for {
	// 	tok, _ := p.scanIgnoreWhitespace()
	// 	if tok == tokenCOLON { // end of inputs
	// 		break
	// 	} else {
	// 		p.unscan() // TODO dont indent?
	// 	}

	// 	tok, lit := p.scanIgnoreWhitespace()
	// 	if tok != tokenID {
	// 		return nil, Error(lit, "identifier")
	// 	}
	// 	in = append(in, lit)
	// }

	// tmp, err := p.nodeRegistry.New(function)
	// if err != nil {
	// 	return nil, err
	// }
	// spec := tmp.(core.INode)
	// node := core.NewCondition(core.TypeAction, function, in, out)
	// spec.Initialize(node)
	// return spec, nil
	return nil, nil
}
