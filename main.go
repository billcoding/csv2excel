package main

import (
	"fmt"
	"github.com/billcoding/csv2excel/cmd"
)

func main() {
	defer func() {
		if re := recover(); re != nil {
			fmt.Println(re)
		}
	}()
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
