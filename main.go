package main

import (
	"fmt"
	"log"
	"os"
)

// main. All the errors need to be handled in this principal main function
func main() {

	// In this part the "challenge.in" file is opened
	file, err := os.Open("challenge.in")
	if err != nil {
		log.Println(err)
	}

	// The file is sended to the scanningFile function
	allConfig, err := scanningFile(file)
	if err != nil {
		log.Println(err)
	}

	// Close the file till the program is done with "defer"
	defer file.Close()

	// Send the response of "scanningFile" to the createFile function
	err = createFile(allConfig)
	if err != nil {
		log.Println(err)
	}

	// Finish de program printing this message
	fmt.Println("Done")
}
