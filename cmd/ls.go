package cmd

import (
	"github.com/byroni/gomunk/pkg/gomunk"
	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(ls)
}

var ls = &cobra.Command{
	Use:     "ls",
	Short:   "List objects in bucket",
	Example: "gomunk ls",
	Run: func(cmd *cobra.Command, args []string) {
		handler := gomunk.GoMunk("aws")
		handler.ListFiles()
	},
}
