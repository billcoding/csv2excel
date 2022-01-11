package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tealeg/xlsx"
	_ "github.com/tealeg/xlsx"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	transformCmd = &cobra.Command{
		Use:     "transform",
		Aliases: []string{"t"},
		Short:   "Transform CSV files to Excel files",
		Long:    "Transform CSV files to Excel files",
		Example: `csv2excel transform -i /to/path -o /to/path`,
		Run:     transform,
	}
	input  string
	output string
)

func init() {
	transformCmd.PersistentFlags().StringVarP(&input, "input", "i", "", "Input directory, default $PWD")
	transformCmd.PersistentFlags().StringVarP(&output, "output", "o", "output", "Output directory")
	rootCmd.AddCommand(transformCmd)
}

func transform(_ *cobra.Command, _ []string) {
	inputDir, _ := filepath.Abs(input)
	outputDir, _ := filepath.Abs(output)
	_ = os.MkdirAll(outputDir, 0700)
	csvFiles, csvNames := readCSVFile(inputDir)

	_, _ = fmt.Printf("Found files: %d\n", len(csvFiles))

	if len(csvFiles) <= 0 {
		return
	}

	for i, f := range csvFiles {
		bytes, _ := ioutil.ReadFile(f)
		ss := strings.Split(string(bytes), "\n")
		ls := make([][]string, 0)
		for _, s := range ss {
			ls = append(ls, strings.Split(s, ";"))
		}
		fileName := csvNames[i]
		idx := strings.LastIndexByte(fileName, '.')
		if idx != -1 {
			oFile := filepath.Join(outputDir, fileName[:idx]+".xlsx")
			_ = writeXlsx(oFile, ls)
			_, _ = fmt.Println(fmt.Sprintf("Transform file => %s", oFile))
		}
	}
}

func readCSVFile(inputDir string) ([]string, []string) {
	files, _ := ioutil.ReadDir(inputDir)
	csvFiles := make([]string, 0)
	csvNames := make([]string, 0)
	for _, f := range files {
		if strings.HasSuffix(strings.ToLower(f.Name()), ".csv") {
			fileName := filepath.Join(inputDir, f.Name())
			csvFiles = append(csvFiles, fileName)
			csvNames = append(csvNames, f.Name())
			_, _ = fmt.Println(fmt.Sprintf("Found file <= %s", fileName))
		}
	}
	return csvFiles, csvNames
}

func writeXlsx(distFile string, data [][]string) error {
	f := xlsx.NewFile()
	sheet, err := f.AddSheet("Sheet1")
	if err != nil {
		return err
	}

	for _, row := range data {
		r := sheet.AddRow()
		for _, cell := range row {
			r.AddCell().SetString(strings.TrimSpace(cell))
		}
	}

	f.Save(distFile)
	return nil
}
