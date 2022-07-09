// Lumeo Package Manager

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func execute(arguments []string) {
	//fmt.Println(len(arguments), arguments)

	switch arguments[0] {
	case "help": // Help command
		fmt.Println("Lumeo Package Manager")
		switch runtime.GOOS {
		case "windows":
			fmt.Println("Windows NT")
		}
	case "install": // Install package
		var installPackage string = arguments[1]
		fmt.Println("Searching for ", installPackage)
	case "create": // Create package
		var createPackageType string = arguments[1]
		var createPackageName string = arguments[2]
		fmt.Println("Creating ", createPackageType, " package ", createPackageName)
		err := os.Mkdir(createPackageName, 0755)
		if err != nil {
			print_error("Error when creating package folder", false)
			log.Fatal(err)
		} else {
			fmt.Println("Package folder created")
		}

		if runtime.GOOS == "windows" {
			fmt.Println("Windows NT")

			cmd := exec.Command("")
			err = cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println("Skipping step for WINDOWS NT only")
		}

	}
}

func print_error(message string, leave bool) {
	fmt.Println("ERROR:  ", message)
	if leave {
		os.Exit(1)
	}
}

func main() {

	arguments := []string{}
	arguments = os.Args[1:]
	//lenght = len(arr)

	if len(arguments) == 0 {
		arguments = append(arguments, "help")
		execute(arguments)
	} else {
		execute(arguments)
	}

}
