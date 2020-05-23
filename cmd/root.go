package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "gomunk",
	Short: "GoMunk is a personal file store and file transfer utility",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gomunk command")
	},
}

func Execute() {
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
