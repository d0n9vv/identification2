package main

import (
	"bufio"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/d0n9vv/identify2/identify"
	"github.com/d0n9vv/identify2/rule"
)

const rootDir = "/home/d0n9vv/code/go_workspace/identification2"

func main() {

	resultFile, errWrite := os.OpenFile(filepath.Join(rootDir, "output/output.txt"), os.O_WRONLY|os.O_CREATE, 0644)
	if errWrite != nil {
		panic(errWrite)
	}
	defer func() { _ = resultFile.Close() }()
	writer := bufio.NewWriter(resultFile)

	phoneNumber := new(rule.DualMode)
	phoneNumber.SetSymbol("PhoneNumber")
	phoneNumber.SetPrefixFilePath(filepath.Join(rootDir, "rulelist/phoneList.txt"))
	phoneNumber.SetSuffix(`([-\s]*\d){8}`)

	idCard := new(rule.SingleMode)
	idCard.SetSymbol("IndentificationCard")
	idCard.SetReStr(`[789][0-9]\d{2}\d{2}(?:1949|19[5-9][0-9]|20[0-1][0-9]|202[0-3])(?:1[3-9]|2[0-4])(?:0[1-9]|[1-2][0-9]|3[0-1])\d{2}\d[X|0-9]`)

	bankCard16 := new(rule.DualMode)
	bankCard16.SetSymbol("BankCard")
	bankCard16.SetPrefixFilePath(filepath.Join(rootDir, "rulelist/bankList.txt"))
	bankCard16.SetSuffix(`([-\s]*\d){10}`)

	bankCard19 := new(rule.DualMode)
	bankCard19.SetSymbol("BankCard")
	bankCard19.SetPrefixFilePath(filepath.Join(rootDir, "rulelist/bankList.txt"))
	bankCard19.SetSuffix(`([-\s]*\d){13}`)

	modes := []rule.Mode{phoneNumber, idCard, bankCard16, bankCard19}

	errRead := filepath.Walk(filepath.Join(rootDir, "testdata"), func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			if filepath.Ext(path) == ".txt" {
				txt := new(identify.Txt)
				txt.SetPath(path)
				txt.SetFileInfo(info)
				txt.Find(writer, modes)
			} else if filepath.Ext(path) == ".csv" {
				csv := new(identify.Csv)
				csv.SetPath(path)
				csv.SetFileInfo(info)
				csv.Find(writer, modes)
			}
		}
		return nil
	})
	if errRead != nil {
		panic(errRead)
	}
}
