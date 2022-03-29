package main

import (
	"fmt"
)

//Insert variables declarations here based on activity
var (
	text             string  = "The following is the account information"
	firstname        string  = "Luke"
	lastname         string  = "SkyWalker"
	age              int     = 20
	weight           float32 = 73.0
	height           float32 = 1.72
	remainingCredits float32 = 123.55
	accountName      string  = "admin"
	accountPassword  string  = "password"
	user             bool    = true
)

func tellMeTypes() {
	fmt.Printf("%T %T %T %T %T %T %T %T %T %T", text, firstname, lastname, age, weight, height, remainingCredits, accountName, accountPassword, user)

}

func main() {
	tellMeTypes()
}
