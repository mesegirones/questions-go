/*
Copyright © 2024 MISERICORDIA GIRONES <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var userId string

// seeStatisticsCmd represents the seeStatistics command
var seeStatisticsCmd = &cobra.Command{
	Use:   "seeStatistics",
	Short: "See user's statistics",
	Long: `See user's statistics. 
	You must provide wich userId you are requesting statistics for using the flag: 
	-u, --userId   User's ID for statistics.  

	If data not provided, the statistics are requested for default user.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if userId == "" {
			log.Fatalln("Missing userId")
			return
		}
		resp, err := http.Get(fmt.Sprintf("http://localhost:8080/question/statistics/%s", userId))
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
	rootCmd.AddCommand(seeStatisticsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// seeStatisticsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// seeStatisticsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	submitAnswersCmd.Flags().StringVarP(&userId, "userId", "u", "1", "User's ID for statistics")
}
