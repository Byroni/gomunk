package cmd

import (
	"log"
	"os"

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
		handler, err := gomunk.GoMunk()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		handler.ListFiles()
	},
}
