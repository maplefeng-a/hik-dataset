package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(topCmd)
}

var topCmd = &cobra.Command{
	Use:   "top",
	Short: "Print cluster info",
	Long:  `Print cluster info`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("---------------------------------")
		fmt.Println("node     gpu_free     gpu_all")
		fmt.Println("hik1        4            2")
		fmt.Println("hik2        4            2")
		fmt.Println("---------------------------------")
	},
}
