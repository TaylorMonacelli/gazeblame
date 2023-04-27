package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "json-parser",
		Short: "A JSON parser using Cobra",
		RunE: func(cmd *cobra.Command, args []string) error {
			jsonStr, err := cmd.Flags().GetString("json")
			if err != nil {
				return err
			}

			// Replace newlines with spaces
			jsonStr = strings.ReplaceAll(jsonStr, "\r\n", " ")
			jsonStr = strings.ReplaceAll(jsonStr, "\n", " ")

			var jsonObj interface{}
			err = json.Unmarshal([]byte(jsonStr), &jsonObj)
			if err != nil {
				return err
			}

			fmt.Println(jsonObj)
			return nil
		},
	}

	rootCmd.Flags().String("json", "", "Multiline JSON string")

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
