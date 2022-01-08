package main

import (
	"flag"
	"fmt"
	_ "github.com/tealeg/xlsx"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	outputDir = flag.String("output", "xlsx-files", "Output XLSX file Directory")
)

func main() {
	flag.Parse()

	_ = os.MkdirAll(*outputDir, 0700)

	root, _ := os.Getwd()
	files, _ := ioutil.ReadDir(root)
	csvFiles := make([]string, 0)
	for _, f := range files {
		//csv
		if strings.HasSuffix(strings.ToLower(f.Name()), ".csv") {
			csvFiles = append(csvFiles, f.Name())
		}
	}
	_, _ = fmt.Printf("find csv files: %d\n", len(csvFiles))

	if len(csvFiles) <= 0 {
		return
	}

	for _, f := range csvFiles {
		bytes, _ := ioutil.ReadFile(f)
		ss := strings.Split(string(bytes), "\n")
		ls := make([][]string, 0)
		for _, s := range ss {
			ls = append(ls, strings.Split(s, ";"))
		}
		idx := strings.LastIndexByte(f, '.')
		if idx != -1 {
			_ = writeXlsx(filepath.Join(*outputDir, f[:idx]+".xlsx"), ls)
		}
	}

}
