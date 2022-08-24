package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// GetEntireJson takes in the name of the json file you want to take
// Opens the json, converts it to bytes, and unmarshales it in a new allStat instance called curSnapShot
// Also checks curSnapShot to see if the json file is empty and stores it in 'exists'
// return curSnapShot and exists
func GetEntireJson(file *os.File) (AllStat, bool) {

	var curSnapShot AllStat

	byteFile, err := ioutil.ReadAll(file)
	json.Unmarshal(byteFile, &curSnapShot)
	if err != nil {
		log.Fatal("Unable to read json file")
	}

	// Check if there is something in our json file
	exists := true
	if curSnapShot.DoesExist == false {
		exists = false
	}
	return curSnapShot, exists
}

// OpenFile takes in string and returns a file with the same name
func OpenFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	return file
}

// WriteToJson takes in an AllStat instance and a filename
// It marshals and writes to the file
func WriteToJson(content AllStat, filename string) {
	out, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(filename, out, 0644)
	if err != nil {
		panic(err)
	}
}

// SetAction checks if you are on the right user
// tells you arent, otherwise sets action
func SetAction(name, filename string, content AllStat, action int) bool {

	_, index, _ := ContainsUserAndPlace(content.Users, name)

	if action == 2 {
		if content.Users[index].Status == 0 {
			fmt.Println("You've yet to open a ticket with this user")
		}
	} else if content.Users[index].Status == action {
		if action == 1 {
			fmt.Println("You cannot open a ticket, you already have one open")
		} else {
			fmt.Println("You cannot close a ticket, you dont have one open")
		}
	} else if content.CurUser.Name != name {
		fmt.Println("You are not on the right user, to change users use 'set-user'")
		fmt.Printf("Current user is: %s\n", content.CurUser.Name)
	} else {
		content.Users[index].LastAction = time.Now()
		content.Users[index].Status = action
		content.CurUser = content.Users[index]
		WriteToJson(content, filename)
		return true
	}
	return false
}

// ContainsUserAndPlace iterates through a slice of Users searching for a specific User.name
// Returns both whether the User.name is in there, and the index its found at
func ContainsUserAndPlace(s []User, name string) (bool, int, User) {
	for x, y := range s {
		if y.Name == name {
			return true, x, y
		}
	}
	return false, 0, User{}
}
