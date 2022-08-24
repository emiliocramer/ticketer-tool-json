package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var openTicketCmd = &cobra.Command{
	Use: "open-ticket",
	Run: func(cmd *cobra.Command, args []string) {
		openTicket(args[0], "USER_DATA.json")
	},
}

func init() {
	rootCmd.AddCommand(openTicketCmd)
}

// openTicket tries to open a ticket on the user you want
func openTicket(name string, filename string) {
	file := OpenFile(filename)
	content, _ := GetEntireJson(file)

	done := SetAction(name, filename, content, 1)
	if done {
		fmt.Println("You've opened a ticket")
	}
}
