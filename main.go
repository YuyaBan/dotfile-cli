package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func main() {
	prompt := promptui.Select{
		Label: "Select OS",
		Items: []string{"MacOS", "Windows", "Ubuntu", "CentOS"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// fmt.Printf("You choose %q\n", result)

	switch result {
	case "MacOS":
		work.setup_Ubuntu()
		fmt.Printf("You choose %q\n", result)
	case "Windows":
		fmt.Printf("You choose %q\n", result)
	case "Ubuntu":
		fmt.Printf("You choose %q\n", result)
	case "CentOS":
		fmt.Printf("You choose %q\n", result)
	}
}
