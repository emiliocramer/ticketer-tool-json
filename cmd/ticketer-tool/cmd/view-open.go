package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var viewOpenCmd = &cobra.Command{
	Use: "view-open",
	Run: func(cmd *cobra.Command, args []string) {
		viewOpen("USER_DATA.json")
	},
}

func init() {
	rootCmd.AddCommand(viewOpenCmd)
}

func viewOpen(filename string) {
	content, _ := GetEntireJson(OpenFile(filename))

	for _, user := range content.Users {
		if user.Status == 1 {
			fmt.Printf("%s has a ticket open\n", user.Name)
		}
	}
}
