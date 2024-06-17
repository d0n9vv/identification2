package rule

import (
	"regexp"
)

type IDCard struct {
	regEx  *regexp.Regexp
	symbol string
}

func NewIDCard(symbol string) *IDCard {
	regEx := regexp.MustCompile(`(\d{18})|(\d{17}[x|X])`)
	return &IDCard{regEx, symbol}
}

func NewIDCardWithRegEx(symbol string, regExStr string) *IDCard {
	regEx := regexp.MustCompile(regExStr)
	return &IDCard{regEx, symbol}
}

func (idCard *IDCard) SetRegEx(str string) {
	idCard.regEx = regexp.MustCompile(str)
}

func (idCard *IDCard) RegEx() *regexp.Regexp {
	return idCard.regEx
}

func (idCard *IDCard) SetSymbol(symbol string) {
	idCard.symbol = symbol
}

func (idCard *IDCard) Symbol() string {
	return idCard.symbol
}

func (idCard *IDCard) FindAll(line []byte) []string {
}

func (idCard *IDCard) FindAllAndValid(line []byte) []string {
}

// 身份证校验算法
func validate(idCard []byte) bool {
	return true
}

func (idCard *IDCard) String() string {
	return "(" + idCard.symbol + ")"
}
