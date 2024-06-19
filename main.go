package main

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/d0n9vv/identify2/pattern"
)

const ROOT = "/home/d0n9vv/code/go_workspace/identification2"

func main() {

	resultFile, errWrite := os.OpenFile(filepath.Join(ROOT, "output/output.txt"), os.O_WRONLY|os.O_CREATE, 0644)
	if errWrite != nil {
		panic(errWrite)
	}
	defer func() { _ = resultFile.Close() }()
	writer := bufio.NewWriter(resultFile)

	idCard := pattern.NewIDCardWithRegEx("IndentificationCard",
		// `[789][0-9]\d{2}\d{2}(?:1949|19[5-9][0-9]|20[0-1][0-9]|202[0-3])(?:1[3-9]|2[0-4])(?:0[1-9]|[1-2][0-9]|3[0-1])\d{2}\d[X|0-9]`)
		`[789][0-9]\d{2}\d{2}(1949|19[5-9][0-9]|20[0-1][0-9]|202[0-3])(1[3-9]|2[0-4])(0[1-9]|[1-2][0-9]|3[0-1])\d{2}\d[X|0-9]`)
	fmt.Println(idCard)

	phoneNumber := pattern.NewPhoneNumber("PhoneNumber")
	phoneNumber.SetRegEx(phoneNumber.PrefixFromFile(filepath.Join(ROOT, "rulelist/phoneList.txt")) + `([-\s]*\d){8}`)
	fmt.Println(phoneNumber)

	bankCard19 := pattern.NewBankCard("BankCard")
	bankCard19.SetRegEx(phoneNumber.PrefixFromFile(filepath.Join(ROOT, "rulelist/bankList.txt")) + `([-\s]*\d){13}`)
	fmt.Println(bankCard19)

	errRead := filepath.Walk(filepath.Join(ROOT, "testdata"), func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			fmt.Println("file:", path)

			file, err := os.Open(path)
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

				idCardFound := idCard.FindAllAndValid(line)
				for _, found := range idCardFound {
					writer.WriteString(fmt.Sprintf("%s-%d-%s-%s\n", info.Name(), row, idCard.Symbol(), found))
				}

				phoneNumberFound := phoneNumber.FindAll(line)
				for _, found := range phoneNumberFound {
					writer.WriteString(fmt.Sprintf("%s-%d-%s-%s\n", info.Name(), row, phoneNumber.Symbol(), found))
				}

				bankCard19Found := bankCard19.FindAllAndValbank(line)
				for _, found := range bankCard19Found {
					writer.WriteString(fmt.Sprintf("%s-%d-%s-%s\n", info.Name(), row, bankCard19.Symbol(), found))
				}

				writer.Flush()
			}
		}
		return nil
	})
	if errRead != nil {
		panic(errRead)
	}
}
