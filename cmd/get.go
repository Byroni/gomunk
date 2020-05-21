package cmd

import (
	"github.com/byroni/gomunk/pkg/gomunk"
	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(get)
}

var get = &cobra.Command{
	Use:     "get",
	Short:   "Get an object by key",
	Example: "gomunk get",
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]

		fileStore := gomunk.AWSFileStore()

		fileStore.GetFile(key)
	},
}
