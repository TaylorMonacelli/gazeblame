package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var myarg string

	rootCmd := &cobra.Command{
		Use:   "gazeblame",
		Short: "A brief description of your application",
		Long:  "A longer description of your application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(string(myarg))
		},
	}

	rootCmd.PersistentFlags().StringVar(&myarg, "jsblob", "", "A multiline string argument")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
