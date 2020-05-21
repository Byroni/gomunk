package cmd

import (
	"errors"

	"github.com/byroni/gomunk/pkg/gomunk"

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
		if len(args) < 1 {
			return errors.New("missing file path\n")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		for _, v := range args {
			handler := gomunk.GoMunk("aws")

			handler.UploadFile(v)
		}
	},
}
