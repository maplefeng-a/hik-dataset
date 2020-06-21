package cmd

import (
	"bytes"
	"fmt"
	"os"

	"github.com/maplefeng-a/hikctl/pkg/exp"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:       "get",
	Short:     "Get object",
	Long:      `Get object`,
	Args:      cobra.OnlyValidArgs,
	ValidArgs: validArgsGet,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("only one arg can be defined")
			os.Exit(1)
		} else {
			if args[0] == "exp" {
				body := exp.Get("ddd")
				buf := new(bytes.Buffer)
				buf.ReadFrom(body)
				s := buf.String()
				fmt.Println(s)
			}
		}
	},
}
