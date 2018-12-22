package lang

import (
	"github.com/alexanderskafte/behaviortree/core"
)

func (p *Parser) parseAction(name string) (core.INode, error) {

	if tok, lit := p.scanIgnoreWhitespace(); tok != tokenPL {
		return nil, Error(lit, "(")
	}

	in, err := p.parseList(tokenCOLON)
	if err != nil {
		return nil, err
	}

	out, err := p.parseList(tokenPR)
	if err != nil {
		return nil, err
	}

	base := core.NewAction(name /* TODO */, in, out)

	return p.nodeRegistry.NewAction(name, base)
}

func (p *Parser) parseList(brk Token) ([]string, error) {
	list := []string{}
	for {
		tok, _ := p.scanIgnoreWhitespace()
		if tok == brk {
			break
		} else {
			p.unscan()
		}

		tok, lit := p.scanIgnoreWhitespace()
		if tok != tokenID {
			return nil, Error(lit, "identifier")
		}
		list = append(list, lit)
	}
	return list, nil
}

// func (p *Parser) parseInput() ([]string, error) {
// 	in := []string{}
// 	for {
// 		tok, _ := p.scanIgnoreWhitespace()
// 		if tok == tokenCOLON { // end of inputs
// 			break
// 		} else {
// 			p.unscan() // TODO dont indent?
// 		}

// 		tok, lit := p.scanIgnoreWhitespace()
// 		if tok != tokenID {
// 			return nil, Error(lit, "identifier")
// 		}
// 		in = append(in, lit)
// 	}
// 	return in, nil
// }

// func (p *Parser) parseOutput() ([]string, error) {
// 	out := []string{}

// 	// Parse the output list
// 	for {
// 		tok, _ := p.scanIgnoreWhitespace()
// 		if tok == tokenPR { // end of outputs
// 			break
// 		} else {
// 			p.unscan() // TODO dont indent?
// 		}

// 		tok, lit := p.scanIgnoreWhitespace()
// 		if tok != tokenID {
// 			return nil, Error(lit, "identifier")
// 		}
// 		out = append(out, lit)
// 	}
// 	return out, nil
// }
