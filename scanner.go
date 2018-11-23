package behaviortree

/*
	CREDIT:
	https://blog.gopheracademy.com/advent-2014/parsers-lexers/
*/

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"unicode"
)

type Token int

const (
	tINVALID Token = iota
	tEOF
	tWS
	tBL
	tBR
	tCOMMA
	tID

	tSEQUENCE
	tSELECTOR
	tACTION
	tINVERTER
)

var lit2tok = map[string]Token{
	"SEQUENCE": tSEQUENCE,
	"SELECTOR": tSELECTOR,
	"ACTION":   tACTION,
	"INVERTER": tINVERTER,
}

const (
	eof   = rune(0)
	bl    = rune('{')
	br    = rune('}')
	comma = rune(',')
)

func isKeyword(lit string) bool {
	_, ok := lit2tok[lit]
	return ok
}

func isWhitespace(ch rune) bool {
	return unicode.IsSpace(ch)
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch)
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
}

func (s *Scanner) Scan() (tok Token, lit string) {
	ch := s.read()

	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()

	} else if isLetter(ch) {
		s.unread()
		return s.scanWord()
	}

	switch ch {
	case eof:
		return tEOF, ""
	case bl:
		return tBL, string(ch)
	case br:
		return tBR, string(ch)
	case comma:
		return tCOMMA, string(ch)
	}

	return tINVALID, string(ch)
}

func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return tWS, buf.String()
}

func (s *Scanner) scanWord() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !(isLetter(ch) || isDigit(ch) || ch == '_' || ch == '.') {
			// Found non-word character
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// If the string matches a keyword then return that keyword.
	lit = strings.ToUpper(buf.String())
	if tok, ok := lit2tok[lit]; ok {
		return tok, lit
	}

	// Otherwise return as a regular identifier.
	return tID, buf.String()
}
