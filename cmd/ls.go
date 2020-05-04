package cmd

import (
	"github.com/byroni/gomunk/pkg/gm"
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
		gm.ListObjects()
	},
}
