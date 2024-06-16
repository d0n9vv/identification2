package rule

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type DualMode struct {
	symbol         string
	prefixFilePath string
	suffix         string
}

func (m *DualMode) SetSymbol(symbol string) {
	m.symbol = symbol
}

func (m *DualMode) Symbol() string {
	return m.symbol
}

func (m *DualMode) SetPrefixFilePath(path string) {
	m.prefixFilePath = path
}

func (m *DualMode) SetSuffix(suf string) {
	m.suffix = suf
}

func (m *DualMode) RegEx() regexp.Regexp {
	var prefix []string

	file, err := os.Open(m.prefixFilePath)
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

	reStr := fmt.Sprintf(`(%s)%s`, strings.Join(prefix, "|"), m.suffix)

	return *regexp.MustCompile(reStr)

}
