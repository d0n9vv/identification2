package pattern

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type Multiple struct {
	regEx  *regexp.Regexp
	symbol string
}

func NewMultiple(symbol string) *Multiple {
	regEx := regexp.MustCompile(`\d{11}`)
	return &Multiple{regEx, symbol}
}

func NewMultipleWithRegEx(symbol string, regExStr string) *Multiple {
	regEx := regexp.MustCompile(regExStr)
	return &Multiple{regEx, symbol}
}

func (mult *Multiple) SetRegEx(str string) {
	mult.regEx = regexp.MustCompile(str)
}

func (mult *Multiple) RegEx() *regexp.Regexp {
	return mult.regEx
}

func (mult *Multiple) SetSymbol(symbol string) {
	mult.symbol = symbol
}

func (mult *Multiple) Symbol() string {
	return mult.symbol
}

func (mult *Multiple) FindAll(line []byte) []string {
	found := []string{}
	replaceRe := regexp.MustCompile(`[-\s]`)

	foundInLine := mult.regEx.FindAll(line, -1)
	for _, item := range foundInLine {
		found = append(found, string(replaceRe.ReplaceAll(item, []byte{})))
	}
	return found
}

func (mult *Multiple) String() string {
	return "{" + mult.symbol + ": `" + mult.regEx.String() + "`}"
}

func (mult *Multiple) PrefixFromFile(path string) string {
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
