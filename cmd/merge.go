package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	mergeCmd = &cobra.Command{
		Use:     "merge",
		Aliases: []string{"m"},
		Short:   "Merge CSV files to one Excel file",
		Long:    "Merge CSV files to one Excel file",
		Example: `csv2excel merge -i /to/path -o res.xlsx`,
		Run:     merge,
	}
	outputFile string
	skipHeader bool
)

func init() {
	mergeCmd.PersistentFlags().StringVarP(&input, "input", "i", "", "Input directory, default $PWD")
	mergeCmd.PersistentFlags().StringVarP(&outputFile, "output", "o", "res.xlsx", "Output file")
	rootCmd.AddCommand(mergeCmd)
}

func merge(_ *cobra.Command, _ []string) {
	inputDir, _ := filepath.Abs(input)
	_ = os.MkdirAll(filepath.Dir(outputFile), 0700)
	csvFiles, csvNames := readCSVFile(inputDir)
	_, _ = fmt.Printf("Found files: %d\n", len(csvFiles))

	if len(csvFiles) <= 0 {
		return
	}

	ls := make([][]string, 0)
	for i, f := range csvFiles {
		bytes, _ := ioutil.ReadFile(f)
		ss := strings.Split(string(bytes), "\n")
		fileName := csvNames[i]
		if idx := strings.LastIndexByte(fileName, '.'); idx != -1 {
			for _, s := range map[bool][]string{true: ss, false: ss[1:]}[i == 0] {
				sss := make([]string, 0)
				sss = append(sss, fileName[:idx])
				sss = append(sss, strings.Split(s, ";")...)
				ls = append(ls, sss)
			}
		}
	}
	_ = writeXlsx(outputFile, ls)
	_, _ = fmt.Println(fmt.Sprintf("Merge file => %s", outputFile))
}
