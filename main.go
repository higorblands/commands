package main

import (
	"log"
	"os"
	"os/exec"
)

func executeCommand(command string, args_array []string) (err error) {

	args := args_array

	commandObject := exec.Command(command, args...)
	commandObject.Stdout = os.Stdout
	commandObject.Stderr = os.Stderr
	err = commandObject.Run()

	if err != nil {

		log.Fatal(err)
		return err
	}

	return nil
}

func main() {

	command := "ls"
	executeCommand(command, []string{"-l"})

}
