package rule

import (
	"regexp"
)

type SingleMode struct {
	symbol string
	reStr  string
}

func (m *SingleMode) SetSymbol(symbol string) {
	m.symbol = symbol
}

func (m *SingleMode) Symbol() string {
	return m.symbol
}

func (m *SingleMode) SetReStr(suf string) {
	m.reStr = suf
}

func (m *SingleMode) RegEx() regexp.Regexp {

	return *regexp.MustCompile(m.reStr)

}
