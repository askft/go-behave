package gbl

import (
	"fmt"

	"github.com/askft/go-behave/core"
)

func (p *Parser) parseExpr() (core.Node, error) {
	p.level++
	defer func() { p.level-- }()

	tok, lit := p.scanIgnoreWhitespace()
	if !isKeyword(lit) {
		return nil, expectError(lit, "keyword")
	}

	it, err := p.accept(tokenIdentifier)
	if err != nil {
		return nil, err
	}
	name := it.lit

	switch tok {
	case tokenComposite:
		return p.parseComposite(name)
	case tokenDecorator:
		return p.parseDecorator(name)
	case tokenCondition:
		return p.parseCondition(name)
	case tokenAction:
		return p.parseAction(name)
	}
	return nil, fmt.Errorf("invalid token")
}

func (p *Parser) parseAssignmentList(brk Token) (map[string]string, error) {
	m := map[string]string{}
	for {
		tok, _ := p.scanIgnoreWhitespace()
		if tok == brk {
			break
		} else {
			p.unscan()
		}

		ass, err := p.parseAssignment()
		if err != nil {
			return nil, err
		}
		m[ass.lhs] = ass.rhs

		t, _ := p.scanIgnoreWhitespace()
		switch t {
		case tokenComma:
			continue
		case brk:
			p.unscan()
		default:
			return nil, fmt.Errorf("balabla")
		}
	}
	return m, nil
}

type assignment struct {
	lhs, rhs string
}

func (p *Parser) parseAssignment() (assignment, error) {
	lhs, err := p.accept(tokenIdentifier)
	if err != nil {
		return assignment{}, err
	}

	if _, err := p.accept(tokenAssign); err != nil {
		return assignment{}, err
	}

	rhs, err := p.accept(tokenLiteral)
	if err != nil {
		return assignment{}, err
	}

	return assignment{lhs.lit, rhs.lit}, nil
}

// func (p *Parser) parseList(brk Token) ([]string, error) {
// 	list := []string{}
// 	for {
// 		if tok, _ := p.scanIgnoreWhitespace(); tok == brk {
// 			break
// 		}
// 		p.unscan()

// 		it, err := p.accept(tokenIdentifier)
// 		if err != nil {
// 			return nil, err
// 		}

// 		list = append(list, it.lit)
// 	}
// 	return list, nil
// }
