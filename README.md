# Introduction

Golang libraries to manage password functions

We use these fucntions in several projects at [Frontware](https://frontware.com) to handle passwords.

- - -

# Functions

## IsCommon

The function retruns trye if the password is in the list of 1 million most used password.
When a user signs up, you can use this function to make sure he is not providing a password that we can easily find in dicitonaries.

The function uses **Bloomfilter** to keep a low memory use.

The list of common passwords comes from https://github.com/danielmiessler/SecLists. [Available here](https://raw.githubusercontent.com/danielmiessler/SecLists/master/Passwords/Common-Credentials/10-million-password-list-top-1000000.txt)

## HashPassword

Returns a bcrypt hash of the password


## NewPassword

Returns a new password.

- - -

# Install

If you use module just import ```github.com/frontware/pwd``` in your code

Else you run:

> go get github.com/frontware/pwd

- - -

# Example

## Sample code

```golang
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
	fmt.Println()
	fmt.Println()
	fmt.Println("Common passwords")
	fmt.Println()
	fmt.Println("Password fjdslkjflkd\t", pwd.IsCommon("fjdslkjflkd"))
	fmt.Println("Password qwerty\t\t", pwd.IsCommon("qwerty"))
}
```

## Results

```
Password                 5EZoCLds
Hash                     $2a$10$yoijtlT0xosZcTd.XWOLUe04zUiNrJj1TAZZSAxJC1zQX/lL.yrhG
Common password

Password fjdslkjflkd     false
Password qwerty          true
```

## On Go Playground

[Try here](https://play.golang.org/p/bRyHXodyT3G)

- - -

# Doc

[Go package doc is here](https://pkg.go.dev/github.com/frontware/pwd)