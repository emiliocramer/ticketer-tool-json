package cmd

import (
	"github.com/spf13/cobra"
	"time"
)

var addUserCmd = &cobra.Command{
	Use: "add-user",
	Run: func(cmd *cobra.Command, args []string) {
		addUser(args[0])
	},
}

func init() {
	rootCmd.AddCommand(addUserCmd)
}

// addUser adds a user
func addUser(name string) {

	file := OpenFile("USER_DATA.json")

	content, exists := GetEntireJson(file)
	if exists == false {
		initialJsonFill(name, "USER_DATA.json")
	} else {
		makeUser(name).loadUser("USER_DATA.json", content)
	}
}

// makeUser will take in a string input and create a User struct with that name
// It will return that User (struct, unmarshalled)
// Turn name(string) into name(type user)
func makeUser(name string) User {
	user := User{
		Name:       name,
		Status:     0,
		LastAction: time.Now(),
	}

	return user
}

// initialJsonFill creates a new AllStat instance, fills it
// Marshals the instance to bytes and writes it to the file
func initialJsonFill(name string, filename string) {

	content := AllStat{}

	content.Users = []User{makeUser(name)}
	content.UserCount = 1
	content.CurUser = content.Users[0]
	content.DoesExist = true

	WriteToJson(content, filename)
}

// loadUser is a method called by a user
// It adds itself to the instance of AllStat it takes in and marshals/writes to the filename it takes in
func (u User) loadUser(filename string, content AllStat) {
	content.Users = append(content.Users, u)
	content.UserCount++
	WriteToJson(content, filename)
}
