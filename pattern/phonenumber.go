package pattern

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type PhoneNumber struct {
	regEx  *regexp.Regexp
	symbol string
}

func NewPhoneNumber(symbol string) *PhoneNumber {
	regEx := regexp.MustCompile(`\d{11}`)
	return &PhoneNumber{regEx, symbol}
}

func NewPhoneNumberWithRegEx(symbol string, regExStr string) *PhoneNumber {
	regEx := regexp.MustCompile(regExStr)
	return &PhoneNumber{regEx, symbol}
}

func (pn *PhoneNumber) SetRegEx(str string) {
	pn.regEx = regexp.MustCompile(str)
}

func (pn *PhoneNumber) RegEx() *regexp.Regexp {
	return pn.regEx
}

func (pn *PhoneNumber) SetSymbol(symbol string) {
	pn.symbol = symbol
}

func (pn *PhoneNumber) Symbol() string {
	return pn.symbol
}

func (pn *PhoneNumber) FindAll(line []byte) []string {
	found := []string{}
	replaceRe := regexp.MustCompile(`[-\s]`)

	foundInLine := pn.regEx.FindAll(line, -1)
	for _, item := range foundInLine {
		found = append(found, string(replaceRe.ReplaceAll(item, []byte{})))
	}
	return found
}

func (pn *PhoneNumber) PrefixFromFile(path string) string {
	var prefix []string

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func() { _ = file.Close() }()

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF { // 用于判断文件是否读取到结尾
			break
		}
		if err != nil {
			panic(err)
		}
		prefix = append(prefix, strings.TrimSpace(string(line)))
	}

	return fmt.Sprintf(`(%s)`, strings.Join(prefix, "|"))
}

func (pn *PhoneNumber) String() string {
	return "{" + pn.symbol + ": `" + pn.regEx.String() + "`}"
}
