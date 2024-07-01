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
	"strings"

	"github.com/spf13/cobra"
)

var bodyInput string

// submitAnswersCmd represents the submitAnswers command
var submitAnswersCmd = &cobra.Command{
	Use:   "submitAnswers",
	Short: "Submit answers for questions",
	Long: `Submit answers for questions that you can see with the command questionList. 
	
	You must provide answers data using the flag: 
	-b, --body   User's answers input.  

	The input must be in the following model: 
	[
		{
			"questionId": "",
			"submittedAnswerID": ""
		}
	]
	There's a dummy data loaded by default. 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		bodyReader := strings.NewReader(bodyInput)
		resp, err := http.Post("http://localhost:8080/question/answers", "application/json", bodyReader)
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

var dummyBody = `[
    {
    "questionId": "1", 
    "submittedAnswerID": "1"
},
{
    "questionId": "2", 
    "submittedAnswerID": "1"
},
{
    "questionId": "3", 
    "submittedAnswerID": "1"
},
{
    "questionId": "4", 
    "submittedAnswerID": "1"
}
]`

func init() {
	rootCmd.AddCommand(submitAnswersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// submitAnswersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// submitAnswersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	submitAnswersCmd.Flags().StringVarP(&bodyInput, "body", "b", dummyBody, "User's answers input")
}
