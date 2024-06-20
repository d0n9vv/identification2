package pattern

import (
	"regexp"
	"strconv"
)

type BankCard struct {
	regEx  *regexp.Regexp
	symbol string
}

func NewBankCard(symbol string) *BankCard {
	regEx := regexp.MustCompile(`(\d{19})`)
	return &BankCard{regEx, symbol}
}

func NewBankCardWithRegEx(symbol string, regExStr string) *BankCard {
	regEx := regexp.MustCompile(regExStr)
	return &BankCard{regEx, symbol}
}

func (bankCard *BankCard) SetRegEx(str string) {
	bankCard.regEx = regexp.MustCompile(str)
}

func (bankCard *BankCard) RegEx() *regexp.Regexp {
	return bankCard.regEx
}

func (bankCard *BankCard) SetSymbol(symbol string) {
	bankCard.symbol = symbol
}

func (bankCard *BankCard) Symbol() string {
	return bankCard.symbol
}

func (bankCard *BankCard) FindAll(line []byte) []string {
	found := []string{}
	replaceRe := regexp.MustCompile(`[-\s]`)

	foundInLine := bankCard.regEx.FindAll(line, -1)
	for _, item := range foundInLine {
		found = append(found, string(replaceRe.ReplaceAll(item, []byte{})))
	}
	return found
}

func (bankCard *BankCard) FindAllAndValid(line []byte) []string {
	found := []string{}
	replaceRe := regexp.MustCompile(`[-\s]`)

	foundInLine := bankCard.regEx.FindAll(line, -1)
	for _, item := range foundInLine {
		if ValidBankCard(item) {
			found = append(found, string(replaceRe.ReplaceAll(item, []byte{})))
		}
	}
	return found
}

// 银行卡号校验算法
func ValidBankCard(item []byte) bool {
	sum := 0
	revPos := 1
	for pos := len(item) - 1; pos >= 0; pos-- {
		if revPos%2 == 0 {
			// 偶数位 乘 2, 再转换成字符串
			x, _ := strconv.Atoi(string(item[pos]))
			strX := []byte(strconv.Itoa(x * 2))

			// 将上一步转换的字符串拆分成数字并求和
			m := 0
			for _, k := range strX {
				n, _ := strconv.Atoi(string(k))
				m += n
			}
			sum += m
		} else {
			x, _ := strconv.Atoi(string(item[pos]))
			sum += x
		}
		revPos++
	}

	return sum%10 == 0
}

func (bankCard *BankCard) String() string {
	return "{" + bankCard.symbol + ": `" + bankCard.regEx.String() + "`}"
}
