package pattern

import (
	"fmt"
	"regexp"
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

func (bankCard *BankCard) FindAllAndValbank(line []byte) []string {
	found := []string{}
	replaceRe := regexp.MustCompile(`[-\s]`)

	foundInLine := bankCard.regEx.FindAll(line, -1)
	for _, item := range foundInLine {
		fmt.Println("------")
		fmt.Println(string(item))
		if ValidBankCard(item) {
			found = append(found, string(replaceRe.ReplaceAll(item, []byte{})))
		}
		fmt.Println("======")
		fmt.Println(string(item))
	}
	return found
}

// 银行卡号校验算法
func ValidBankCard(item []byte) bool {
	sum := 0
	// slices.Reverse(item)

	for i, v := range item {
		// 	if v >= 48 && v <= 57 { // 0 - 9
		// 		vv, _ := strconv.Atoi(string(v))
		// 		sum += vv * factor[i]
		// 	} else { // X | x
		// 		sum += 10 * factor[i]
		// 	}
	}
	return sum%11 == 1
}

func (bankCard *BankCard) String() string {
	return "{" + bankCard.symbol + ": `" + bankCard.regEx.String() + "`}"
}
