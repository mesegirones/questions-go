/*
Copyright © 2024 MISERICORDIA GIRONES <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// questionListCmd represents the questionList command
var questionListCmd = &cobra.Command{
	Use:   "questionList",
	Short: "Request questions list",
	Long:  `Request questions list`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("http://localhost:8080/question/list")
		if err != nil {
			log.Fatalln(err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		var prettyJSON bytes.Buffer
		json.Indent(&prettyJSON, body, "", "\t")
		log.Println(prettyJSON.String())
	},
}

func init() {
	rootCmd.AddCommand(questionListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// questionListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// questionListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
