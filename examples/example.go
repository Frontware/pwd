package main

import (
	"fmt"

	"github.com/frontware/pwd"
)

func main() {

	// Generate new password of 8 chars
	pass := pwd.NewPassword(8)
	// Get hash
	hash, _ := pwd.HashPassword(pass)

	fmt.Println("Password\t\t", pass)
	fmt.Println("Hash\t\t\t", hash)

	fmt.Println("Common password")
	fmt.Println()
	fmt.Println("Password fjdslkjflkd\t", pwd.IsCommon("fjdslkjflkd"))
	fmt.Println("Password qwerty\t\t", pwd.IsCommon("qwerty"))

}
