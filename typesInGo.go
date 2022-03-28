package main

import "fmt"

//Insert variables declarations here based on activity
var (
	text      string = "The following is the account information"
	firstname string = "Luke"
	lastname  string = "SkyWalker"
)

func tellMeTypes() {
	fmt.Printf("%T %T %T", text, firstname, lastname)
}

func main() {
	tellMeTypes()
}
