// Lumeo Package Manager

package main

import (
	"fmt"
	"os"
)

func execute(command string) {

}

func print_error(message string) {
	fmt.Println("ERROR:  ", message)
	os.Exit(1)
}

func main() {

	lenght := len(os.Args)
	//var arr []string
	//lenght = len(arr)

	if lenght <= 2 {
		print_error("Not enough arguments passed")
	} else {
		//execute("dsads")
	}

}
