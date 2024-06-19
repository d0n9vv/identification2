package pattern

import (
	"regexp"
	"strconv"
)

type Single struct {
	regEx  *regexp.Regexp
	symbol string
}

func NewSingle(symbol string) *Single {
	regEx := regexp.MustCompile(`(\d{18})|(\d{17}[x|X])`)
	return &Single{regEx, symbol}
}

func NewSingleWithRegEx(symbol string, regExStr string) *Single {
	regEx := regexp.MustCompile(regExStr)
	return &Single{regEx, symbol}
}

func (single *Single) SetRegEx(str string) {
	single.regEx = regexp.MustCompile(str)
}

func (single *Single) RegEx() *regexp.Regexp {
	return single.regEx
}

func (single *Single) SetSymbol(symbol string) {
	single.symbol = symbol
}

func (single *Single) Symbol() string {
	return single.symbol
}

func (single *Single) FindAll(line []byte) []string {
	found := []string{}
	replaceRe := regexp.MustCompile(`[-\s]`)

	foundInLine := single.regEx.FindAll(line, -1)
	for _, item := range foundInLine {
		found = append(found, string(replaceRe.ReplaceAll(item, []byte{})))
	}
	return found
}

func (single *Single) FindAllAndValid(line []byte) []string {
	found := []string{}
	replaceRe := regexp.MustCompile(`[-\s]`)

	foundInLine := single.regEx.FindAll(line, -1)
	for _, item := range foundInLine {
		if single.Validate(item) {
			found = append(found, string(replaceRe.ReplaceAll(item, []byte{})))
		}
	}
	return found
}

// 身份证号码校验算法
func (single *Single) Validate(item []byte) bool {
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

func (single *Single) String() string {
	return "(" + single.symbol + ")"
}
