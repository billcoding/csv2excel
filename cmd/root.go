package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "csv2excel",
		Short:   "csv2excel::CSV转换工具",
		Long:    `csv2excel::CSV转换工具`,
		Version: version,
	}
	version = "1.0.1"
)

func Execute() error {
	return rootCmd.Execute()
}
