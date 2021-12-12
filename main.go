package main

import (
	"fmt"

	"github.com/YuyaBan/dotfile-cli/work"
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

	fmt.Printf("You choose %q\n", result)

	switch result {
	case "MacOS":
		fmt.Printf("but not ready yet.\n", result)
	case "Windows":
		fmt.Printf("but not ready yet.\n", result)
	case "Ubuntu":
		work.SetupUbuntu()
	case "CentOS":
		fmt.Printf("but not ready yet.\n", result)
	}
}
