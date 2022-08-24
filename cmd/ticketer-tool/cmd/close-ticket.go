package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var closeTicketCmd = &cobra.Command{
	Use: "close-ticket",
	Run: func(cmd *cobra.Command, args []string) {
		closeTicket(args[0], "USER_DATA.json")
	},
}

func init() {
	rootCmd.AddCommand(closeTicketCmd)
}

// closeTicket tries to close a ticket on the user you want
func closeTicket(name string, filename string) {
	file := OpenFile(filename)
	content, _ := GetEntireJson(file)

	done := SetAction(name, filename, content, 2)
	if done {
		fmt.Println("You've closed a ticket")
	}
}
