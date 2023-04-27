package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "json-parser",
		Short: "A JSON parser using Cobra",
		RunE: func(cmd *cobra.Command, args []string) error {
			var jsonObj interface{}

			jsonFile, _ := cmd.Flags().GetString("file")
			if jsonFile != "" {
				jsonBytes, err := ioutil.ReadFile(jsonFile)
				if err != nil {
					return err
				}

				err = json.Unmarshal(jsonBytes, &jsonObj)
				if err != nil {
					return err
				}
			} else {
				jsonStr, _ := cmd.Flags().GetString("json")

				err := json.Unmarshal([]byte(jsonStr), &jsonObj)
				if err != nil {
					return err
				}
			}

			fmt.Println(jsonObj)
			return nil
		},
	}

	rootCmd.Flags().String("json", "", "JSON string")
	rootCmd.Flags().String("file", "", "JSON file path")

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
