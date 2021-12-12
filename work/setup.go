package work

import (
	"bytes"
	_ "embed"
	"fmt"
	"os/exec"
)

//go:embed scripts/install.sh
var contents string

func SetupUbuntu() {
	fmt.Printf("[+] start setup Ubuntu\n")

	cmdline := exec.Command("bash")
	var buf = new(bytes.Buffer)
	buf.WriteString(contents)
	cmdline.Stdin = buf

	out, err := cmdline.Output()
	if err != nil {
		fmt.Printf("execCommand faild %v\n", err)
		return
	}

	fmt.Println(string(out))
}
