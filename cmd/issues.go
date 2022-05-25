package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var issuesCmd = &cobra.Command{
	Use:   "issues",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("issues")
	},
}

func init() {
	rootCmd.AddCommand(issuesCmd)
}
