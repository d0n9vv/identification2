package identify

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"
	"regexp"

	"github.com/d0n9vv/identify2/rule"
)

type Txt struct {
	path     string
	fileInfo fs.FileInfo
}

func (txt *Txt) SetPath(path string) {
	txt.path = path
}

func (txt *Txt) SetFileInfo(info fs.FileInfo) {
	txt.fileInfo = info
}

func (txt *Txt) Find(writer *bufio.Writer, modes []rule.Mode) {
	fmt.Println("file:", txt.path)
	// fmt.Println(strings.Repeat("=", 40))

	file, err := os.Open(txt.path)
	if err != nil {
		panic(err)
	}
	defer func() { _ = file.Close() }()

	reader := bufio.NewReader(file)
	for row := 1; ; row++ {
		line, _, err := reader.ReadLine() // 按行读取文件
		if err == io.EOF {                // 用于判断文件是否读取到结尾
			break
		}
		if err != nil {
			panic(err)
		}

		replaceRe := regexp.MustCompile(`[-\s]`)

		for _, mode := range modes {
			re := mode.RegEx()
			found := re.FindAll(line, -1)
			for _, f := range found {
				f = replaceRe.ReplaceAll(f, []byte{})
				writer.WriteString(fmt.Sprintf("%s-%d-%s-%s\n", txt.fileInfo.Name(), row, mode.Symbol(), f))
			}
		}
		writer.Flush()
	}

}
