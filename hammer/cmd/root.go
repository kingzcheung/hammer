package cmd

import (
	"fmt"
	"github.com/kingzcheung/hammer/box"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
)

var namePackage string

const (
	// fmt.Printf("\x1b[%d;%dmhello world \x1b[0m 46: 深绿 31: 红 \n", 46, 31)
	green = "\x1b[32m"
	red   = "\x1b[31m"
)

var force bool

var rootCmd = &cobra.Command{
	Use:   "hammer",
	Short: "Simple tool to embed files in Go binary",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("please enter the directory to embed")
		}
		_, err := os.Stat(args[0])
		if err != nil {
			return err
		}
		// if !stat.IsDir() {
		// 	return errors.Errorf("%s is not directory", args[0])
		// }
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		zbox := box.NewZbox(args[0])
		zbox.SetForceRemove(force)
		if namePackage != "" {
			zbox.SetNamePackage(namePackage)
		}
		err := zbox.Hammer()
		if err != nil {
			fmt.Printf("%s%s\n", red, err.Error())
			return
		}
		fmt.Printf("%sSuccess!\n", green)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&namePackage, "package", "p", "", "The package name")
	rootCmd.Flags().BoolVarP(&force, "force", "f", false, "Forced generation")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
