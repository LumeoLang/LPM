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
	case "welcome": // Welcome message
		fmt.Println("██      ██████  ███    ███ ")
		fmt.Println("██      ██   ██ ████  ████ ")
		fmt.Println("██      ██████  ██ ████ ██ ")
		fmt.Println("██      ██      ██  ██  ██ ")
		fmt.Println("███████ ██      ██      ██ ")
		fmt.Printf("\n")
		fmt.Println("Lumeo Package Manager")
		switch runtime.GOOS {
		case "windows":
			fmt.Println("Windows NT")
		case "linux":
			fmt.Println("Linux")
		}
		fmt.Println("Start by using 'lpm' + command ('lpm help')")

	case "help": // Help command
		fmt.Println("Lumeo Package Manager")
		fmt.Println("Commands:")
		fmt.Println("  help - Show this help")
		fmt.Println("  install - Install a package")
		fmt.Println("  uninstall - Uninstall a package")
		fmt.Println("  list - List all installed packages")
		fmt.Println("  update - Update all installed packages")
		fmt.Println("  create - Creates a new package")

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

	default: // Unknown command
		print_error("Command not found", true)

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
		arguments = append(arguments, "welcome")
		execute(arguments)
	} else {
		execute(arguments)
	}

}

// End of file
