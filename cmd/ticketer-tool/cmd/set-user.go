package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var setUserCmd = &cobra.Command{
	Use: "set-user",
	Run: func(cmd *cobra.Command, args []string) {
		setUser(args[0], "USER_DATA.json")
	},
}

func init() {
	rootCmd.AddCommand(setUserCmd)
}

// setUser gets the whole json file and checks to see if name is valid
// sets to the valid name, prints an error otherwise
func setUser(name string, filename string) {
	file := OpenFile(filename)

	content, _ := GetEntireJson(file)

	contains, index := ContainsUserAndPlace(content.Users, name)
	if contains {
		content.CurUser = content.Users[index]
		out, _ := json.Marshal(content)
		ioutil.WriteFile(filename, out, 0644)
		fmt.Printf("Set user to: %s\n", name)
	} else {
		fmt.Println("User does not exist, if you would like to create a user use the command 'add-user'")
	}

}
