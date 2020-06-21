package cmd

import (
	"bytes"
	"fmt"
	"os"

	"github.com/maplefeng-a/hikctl/pkg/exp"
	"github.com/spf13/cobra"
)

func init() {

	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&file, "file", "f", "", "define file")
}

var deleteCmd = &cobra.Command{
	Use:       "delete",
	Short:     "Delete object",
	Long:      `Delete object`,
	Args:      cobra.OnlyValidArgs,
	ValidArgs: validArgsDelete,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("only one arg can be defined")
			os.Exit(1)
		} else {
			if args[0] == "exp" {
				body := exp.Delete("ddd")
				buf := new(bytes.Buffer)
				buf.ReadFrom(body)
				s := buf.String()
				fmt.Println(s)
			}
		}
	},
}
