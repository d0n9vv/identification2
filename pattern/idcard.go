package pattern

import (
	"regexp"
	"strconv"
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
	found := []string{}
	replaceRe := regexp.MustCompile(`[-\s]`)

	foundInLine := idCard.regEx.FindAll(line, -1)
	for _, item := range foundInLine {
		found = append(found, string(replaceRe.ReplaceAll(item, []byte{})))
	}
	return found
}

func (idCard *IDCard) FindAllAndValid(line []byte) []string {
	found := []string{}
	replaceRe := regexp.MustCompile(`[-\s]`)

	foundInLine := idCard.regEx.FindAll(line, -1)
	for _, item := range foundInLine {
		if ValidIDCard(item) {
			found = append(found, string(replaceRe.ReplaceAll(item, []byte{})))
		}
	}
	return found
}

// 身份证号码校验算法
func ValidIDCard(item []byte) bool {
	if len(item) != 18 {
		return false
	}
	factor := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2, 1}
	sum := 0
	for i, v := range item {
		if v >= 48 && v <= 57 { // 0 - 9
			vv, _ := strconv.Atoi(string(v))
			sum += vv * factor[i]
		} else { // X | x
			sum += 10 * factor[i]
		}
	}
	return sum%11 == 1
}

func (idCard *IDCard) String() string {
	return "{" + idCard.symbol + ": `" + idCard.regEx.String() + "`}"
}
