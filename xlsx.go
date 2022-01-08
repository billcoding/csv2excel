package main

import (
	"github.com/tealeg/xlsx"
)

func writeXlsx(distFile string, data [][]string) error {
	f := xlsx.NewFile()
	sheet, err := f.AddSheet("Sheet1")
	if err != nil {
		return err
	}

	for _, row := range data {
		r := sheet.AddRow()
		for _, cell := range row {
			r.AddCell().SetString(cell)
		}
	}

	f.Save(distFile)
	return nil
}
