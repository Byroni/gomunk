package cmd

import (
	"errors"

	"github.com/byroni/gomunk/pkg/gm"

	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(upload)
}

var upload = &cobra.Command{
	Use:     "upload",
	Short:   "Upload a file to personal s3 bucket",
	Example: "gomunk upload path/to/file.txt",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) <= 1 {
			return errors.New("Missing file path")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		for _, v := range args {
			client := gm.NewGoMunkUpload(v)
			client.UploadFile()
		}
	},
}
