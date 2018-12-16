package lang

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
	tPL
	tPR
	tCOMMA
	tCOLON
	tEMARK
	tQMARK
	tID

	tSEQUENCE
	tSELECTOR
	tACTION
	tINVERTER
)

var lit2tok = map[string]Token{
	"SEQUENCE": tSEQUENCE,
	"SELECTOR": tSELECTOR,
	"INVERTER": tINVERTER,
	"ACTION":   tACTION,
}

const (
	rEOF          = rune(0)
	rBracketLeft  = rune('{')
	rBracketRight = rune('}')
	rParenLeft    = rune('(')
	rParenRight   = rune(')')
	rComma        = rune(',')
	rColon        = rune(':')
	rExclamation  = rune('!')
	rQuestion     = rune('?')
)

// TokenIsEOF returns true if `tok` is an EOF token.
func TokenIsEOF(tok Token) bool {
	return tok == tEOF
}

// TokenIsWhitespace returns true if `tok` is a whitespace token.
func TokenIsWhitespace(tok Token) bool {
	return tok == tWS
}

// TokenIsInvalid returns true if `tok` is an invalid token.
func TokenIsInvalid(tok Token) bool {
	return tok == tINVALID
}

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

// NewScanner returns a scanner that reads from `r`.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return rEOF
	}
	return ch
}

func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
}

// Scan scans one token, and returns the token and the scanned string.
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
	case rEOF:
		return tEOF, ""
	case rBracketLeft:
		return tBL, string(ch)
	case rBracketRight:
		return tBR, string(ch)
	case rParenLeft:
		return tPL, string(ch)
	case rParenRight:
		return tPR, string(ch)
	case rColon:
		return tCOLON, string(ch)
	case rExclamation:
		return tEMARK, string(ch)
	case rQuestion:
		return tQMARK, string(ch)
	case rComma:
		return tCOMMA, string(ch)
	}

	return tINVALID, string(ch)
}

func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())
	for {
		if ch := s.read(); ch == rEOF {
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
		if ch := s.read(); ch == rEOF {
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
