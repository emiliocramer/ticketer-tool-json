package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var longestOpenCmd = &cobra.Command{
	Use: "longest-open",
	Run: func(cmd *cobra.Command, args []string) {
		longestOpen("USER_DATA.json")
	},
}

func init() {
	rootCmd.AddCommand(longestOpenCmd)
}

func longestOpen(filename string) {
	content, _ := GetEntireJson(OpenFile(filename))
	var longestUser User
	for _, user := range content.Users {
		if user.LastAction.After(longestUser.LastAction) {
			longestUser = user
		}
	}
	fmt.Printf("%s has had a ticket open the longest\n", longestUser.Name)
}
