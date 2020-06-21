/*
Copyright 2020 maplefeng-a.
*/

// Package cmd is responsible for dataset manager.
package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

const (
	hikctl = "hikctl"
)

var rootCmd = &cobra.Command{
	Use:   hikctl,
	Short: hikctl + " is a cli of hikmars3 platform.",
	Long:  hikctl + " is a cli of hikmars3 platform.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello cobra!")

	},
}

var (
	cfgFile           string
	validArgsCreate   = []string{"exp"}
	validArgsDelete   = []string{"exp"}
	validArgsGet      = []string{"exp"}
	validArgsDescribe = []string{"exp"}
	file              string
	image             string
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hikctl.yaml)")
	//rootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
	rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".hikctl" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".hikctl")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

// Execute comments here
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
