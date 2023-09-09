package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-vblog/version"
	"os"
)

var (
	vers bool
)

var RootCommand = &cobra.Command{
	Use:   "vblog",
	Short: "vblog backend",
	Long:  "vblog 博客后端管理系统",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println(version.FullVersion())
			return nil
		}
		return cmd.Help()
	},
}

func Execute() {
	if err := RootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCommand.PersistentFlags().BoolVarP(&vers, "version", "v", false, "the vblog api version")
}
