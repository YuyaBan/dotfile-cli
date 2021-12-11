package work

import (
	"fmt"
	"os/exec"
)

func SetupUbuntu() {
	fmt.Printf("[+] start setup Ubuntu\n")

	out, err := exec.Command("echo", "foobar").Output()
	if err != nil {
		fmt.Printf("execCommand faild %v\n", err)
		return
	}

	fmt.Println(string(out))
}
