package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Hello")

	startJava()
}

func startJava() {
	fmt.Println("Starting java...")

	cmd := exec.Command("java", "-jar", "target/mainModule-1.0-SNAPSHOT-jar-with-dependencies.jar")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	}
	cmd.Wait()

	fmt.Println("Started java.")
}
