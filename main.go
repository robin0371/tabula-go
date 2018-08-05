package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	tabula "tabula-go/lib"
)

// CLI for extract data using tabula
func main() {
	log.Println("This is CLI for extract data using tabula")
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	args := make([]string, 0, 15)
	// Add path to tabula jar file
	args = append(args, []string{"-jar", path.Join(pwd, "lib", tabula.TabulaJar)}...)
	// Add args from command line
	args = append(args, os.Args[1:]...)

	cmd := exec.Command("java", args...)
	log.Printf("Command: %s\n", strings.Join(cmd.Args, " "))

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error: %s\nResult: %s", err, output)
	} else {
		log.Printf("Success\nResult: %s", output)
	}
}
