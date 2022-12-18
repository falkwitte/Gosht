package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Set up a map of builtin commands
	builtins := map[string]func([]string){
		"cd":     cd,
		"echo":   echo,
		"export": export,
		"exit":   exit,
	}

	// Set up a slice to store command history
	history := make([]string, 0)

	// Start the main loop
	for {
		// Print the prompt
		fmt.Print("$ ")

		// Read a line of input from the user
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		// Add the input to the command history
		history = append(history, input)

		// Split the input into words
		words := strings.Split(input, " ")

		// Check if the first word is a builtin command
		if f, ok := builtins[words[0]]; ok {
			// If it is, call the builtin function
			f(words[1:])
		} else {
			// If it's not a builtin command, execute it as a subprocess
			cmd := exec.Command(words[0], words[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		}
	}
}

// cd is a builtin function that changes the current working directory
func cd(args []string) {
	if len(args) != 1 {
		fmt.Println("usage: cd directory")
		return
	}
	err := os.Chdir(args[0])
	if err != nil {
		fmt.Println(err)
	}
}

// echo is a builtin function that prints its arguments to standard output
func echo(args []string) {
	fmt.Println(strings.Join(args, " "))
}

// export is a builtin function that sets an environment variable
func export(args []string) {
	if len(args) != 2 {
		fmt.Println("usage: export VARNAME=value")
		return
	}
	parts := strings.Split(args[0], "=")
	if len(parts) != 2 {
		fmt.Println("usage: export VARNAME=value")
		return
	}
	err := os.Setenv(parts[0], parts[1])
	if err != nil {
		fmt.Println(err)
	}
}

// exit is a builtin function that exits the shell
func exit(args []string) {
	os.Exit(0)
}
