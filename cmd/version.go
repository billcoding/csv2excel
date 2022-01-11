package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

var (
	version    = "1.0.1"
	versionCmd = &cobra.Command{
		Use:     "version",
		Aliases: []string{"v", "ver"},
		Short:   "Show csv2excel version",
		Long:    "Show csv2excel version",
		Run: func(cmd *cobra.Command, args []string) {
			_, _ = fmt.Printf("csv2excel %s\n%s\n", version, runtime.Version())
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
