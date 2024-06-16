package identify

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/fs"
	"os"
	"regexp"

	"github.com/d0n9vv/identify2/rule"
)

type Csv struct {
	path     string
	fileInfo fs.FileInfo
}

func (csv *Csv) SetPath(path string) {
	csv.path = path
}

func (csv *Csv) SetFileInfo(info fs.FileInfo) {
	csv.fileInfo = info
}

func (csvFile *Csv) Find(writer *bufio.Writer, modes []rule.Mode) {
	fmt.Println("file:", csvFile.path)
	file, err := os.Open(csvFile.path)
	if err != nil {
		panic(err)
	}
	defer func() { _ = file.Close() }()

	reader := csv.NewReader(file)
	for row := 1; ; row++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		replaceRe := regexp.MustCompile(`[-\s]`)

		phoneNumberRe := modes[0].RegEx()
		idCardRe := modes[1].RegEx()
		bankCard16Re := modes[2].RegEx()
		bankCard19Re := modes[3].RegEx()

		for i, item := range record {
			if i == 1 {
				found := []string{}
				found = append(found, phoneNumberRe.FindAllString(item, -1)...)
				for _, f := range found {
					f = replaceRe.ReplaceAllString(f, "")
					writer.WriteString(fmt.Sprintf("%s-%d-%s-%s\n", csvFile.fileInfo.Name(), row, modes[0].Symbol(), f))
				}

			} else if i == 2 {
				found := []string{}
				found = append(found, idCardRe.FindAllString(item, -1)...)
				for _, f := range found {
					f = replaceRe.ReplaceAllString(f, "")
					writer.WriteString(fmt.Sprintf("%s-%d-%s-%s\n", csvFile.fileInfo.Name(), row, modes[1].Symbol(), f))
				}

			} else if i == 3 {
				found := []string{}
				found = append(found, bankCard16Re.FindAllString(item, -1)...)
				found = append(found, bankCard19Re.FindAllString(item, -1)...)
				for _, f := range found {
					f = replaceRe.ReplaceAllString(f, "")
					writer.WriteString(fmt.Sprintf("%s-%d-%s-%s\n", csvFile.fileInfo.Name(), row, modes[2].Symbol(), f))
				}

			}
		}
		writer.Flush()

	}

}
