package cmd

import (
	"bytes"
	"fmt"
	"os"

	"github.com/maplefeng-a/hikctl/pkg/exp"
	"github.com/spf13/cobra"
)

func init() {

	rootCmd.AddCommand(describeCmd)
}

var describeCmd = &cobra.Command{
	Use:       "describe",
	Short:     "Describe object",
	Long:      `Describe object`,
	Args:      cobra.OnlyValidArgs,
	ValidArgs: validArgsDescribe,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("only one arg can be defined")
			os.Exit(1)
		} else {
			if args[0] == "exp" {
				body := exp.Describe("ddd")
				buf := new(bytes.Buffer)
				buf.ReadFrom(body)
				s := buf.String()
				fmt.Println(s)
			}
		}
	},
}
