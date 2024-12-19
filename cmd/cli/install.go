package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type options struct {
	accessKey string
}

func main() {
	var options options

	flag.StringVar(&options.accessKey, "key", "", "API Key")
	flag.Parse()

	if len(options.accessKey) == 0 {
		log.Fatal("No access key provided")
	}

	os.Setenv("ACCESS_KEY", options.accessKey)

	// download the beacon executable

	// run the executable passing the API as an option

	//get current working directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	//change directory to the binary location
	err = os.Chdir(dir + "/beacon")

	fmt.Println(dir + "/beacon")

	// run the executable passing the API as an option
	// Specify the binary you want to run (e.g., `ls` or any executable file)
	binary := "./beacon"

	// Specify any arguments the binary might require
	args := []string{fmt.Sprintf("-key=%s", options.accessKey)}

	// Create the command
	cmd := exec.Command(binary, args...)

	// Set the output to the standard output (stdout)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running command: %v\n", err)
	} else {
		fmt.Println("Command executed successfully.")
	}
}

func downloadBinary() {

}
