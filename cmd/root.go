package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "csv2excel",
		Short:   "csv2excel::CSV转换工具",
		Long:    `csv2excel::CSV转换工具`,
		Version: "1.0.2",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
