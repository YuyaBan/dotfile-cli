package work

import {
	"fmt"
}

func setup_Ubuntu () {
	fmt.Printf("[+] start setup Ubuntu\n")

	out, err := execCommand("echo", "foobar").Output

	fmt.Println(string(out))
}