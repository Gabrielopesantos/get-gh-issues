package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gghi",
	Short: "Short description",
	Long:  `Long description`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	log.Println("Hello world")
	//},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
