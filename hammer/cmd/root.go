package cmd

import (
	"fmt"
	"github.com/kingzcheung/hammer/box"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
)

var namePackage string

var rootCmd = &cobra.Command{
	Use:   "hammer",
	Short: "Simple tool to embed files in Go binary",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("please enter the directory to embed")
		}
		stat, err := os.Stat(args[0])
		if err != nil {
			return err
		}
		if !stat.IsDir() {
			return errors.Errorf("%s is not directory", args[0])
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		zbox := box.NewZbox(args[0])
		if namePackage != "" {
			zbox.SetNamePackage(namePackage)
		}
		fmt.Println(zbox.Hammer())
	},
}

func init() {
	rootCmd.Flags().StringVarP(&namePackage, "package", "p", "", "The package name")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
