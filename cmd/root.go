package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gghi",
	Short: "Short description",
	Long:  `Long description`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Hello world")
	},
}

var helpCmd = &cobra.Command{
	Use:   "auth",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("auth desc")
	},
}

func Execute() {
	rootCmd.AddCommand(helpCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
