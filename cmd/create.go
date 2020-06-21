package cmd

import (
	"bytes"
	"fmt"
	"os"

	"github.com/maplefeng-a/hikctl/pkg/exp"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&file, "file", "f", "", "define file")
	createCmd.Flags().StringVarP(&image, "image", "i", "", "define image")
}

var createCmd = &cobra.Command{
	Use:       "create",
	Short:     "Create object",
	Long:      `Create object`,
	Args:      cobra.OnlyValidArgs,
	ValidArgs: validArgsCreate,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("only one arg can be defined")
			os.Exit(1)
		} else {
			if args[0] == "exp" {
				body := exp.Create("ddd")
				buf := new(bytes.Buffer)
				buf.ReadFrom(body)
				s := buf.String()
				fmt.Println(s)
			}
		}
	},
}
