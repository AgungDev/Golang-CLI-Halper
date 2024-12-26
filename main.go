package main

import (
	"fmt"
	"strings"

	cligz "github.com/AgungDev/Golang-CLI-Halper/cliGz"
)

func main(){
	mainMenu()
}

func mainMenu(){
	menu := []string {
		"Customer", "Exit",
	}

	menu_choice := cligz.Menu("Enter your choice", menu)
	switch menu_choice {
		case 1:
			customerMenu()
			return
		case 2:
			return // exit
		
	}
}

func customerMenu(){
	customerMenu := []string {
		"Create", "Read", "Update", "Delate","Back To Menu",
	}
	customer_choice := cligz.Menu("Enter your choice", customerMenu)
	switch customer_choice {
		case 1:
			insertCustomer()
			return
		case 2:
			readCustomer()
			return
		case 5:
			mainMenu()
			return 
		default:
			mainMenu()
			return
	}
}

func insertCustomer(){
	// check id
	id := cligz.IntegerResult("Insert Customer ID", "Customer ID Not Founds!")
	form_user := cligz.Form([]string {
		"Nama","Email", "Umur",
	})
	fmt.Println(id, form_user)
	fmt.Println(strings.Repeat("=", 50))
	customerMenu()
}

func readCustomer(){

}