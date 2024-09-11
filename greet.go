package greet

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func GreetUser(stdin io.Reader, stdout io.Writer) {
	name := "Vlad"
	fmt.Fprintln(stdout, "Whats is your name?")
	input := bufio.NewScanner(stdin)
	if input.Scan() {
		name = input.Text()
	}
	fmt.Fprintf(stdout, "Hello, %s\n", name)
}

func Main() int {
	GreetUser(os.Stdin, os.Stdout)
	return 0
}
