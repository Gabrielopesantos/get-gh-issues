package cmd

import (
	"fmt"
	"github.com/gabrielopesantos/get-gh-issues/internal/auth"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "",
	Long:  ``,
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logout")
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("status")
	},
}

func init() {
	rootCmd.AddCommand(authCmd)

	authCmd.AddCommand(auth.NewCmdLogin())
	authCmd.AddCommand(logoutCmd)
	authCmd.AddCommand(statusCmd)
}
