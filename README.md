<img src="https://www.frontware.com/images/img/fw-logo.png" alt="Frontware" width="120"/>

# Introduction



Golang libraries to manage password functions

We use these functions in several projects at [Frontware](https://frontware.com) to handle passwords.



- - -

# Functions

## IsCommon

The function retruns trye if the password is in the list of 1 million most used password.
When a user signs up, you can use this function to make sure he is not providing a password that we can easily find in dicitonaries.

The function uses [**Bloomfilter**](https://en.wikipedia.org/wiki/Bloom_filter#:~:text=A%20Bloom%20filter%20is%20a,a%20member%20of%20a%20set.) to keep a low memory consumption.

The list of common passwords comes from https://github.com/danielmiessler/SecLists.

Get the list of million passwords [here](https://raw.githubusercontent.com/danielmiessler/SecLists/master/Passwords/Common-Credentials/10-million-password-list-top-1000000.txt).

## HashPassword

Returns a **bcrypt** hash of the password as a string.


## NewPassword

Returns a new random password.

It's never a good idea to let user set their own password. This functions generates a random password for your user.

- - -

# Install

If you use module just import ```github.com/frontware/pwd``` in your code.

If you don't use go mod then you download the module:

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

[![Try it on Go playground](https://img.shields.io/static/v1?label=Golang&message=Playground&color=blue)](https://play.golang.org/p/bRyHXodyT3G)

## Results

```
Password                 5EZoCLds
Hash                     $2a$10$yoijtlT0xosZcTd.XWOLUe04zUiNrJj1TAZZSAxJC1zQX/lL.yrhG
Common password

Password fjdslkjflkd     false
Password qwerty          true
```


- - -

# Doc

See the doc online

[![PkgGoDev](https://pkg.go.dev/badge/github.com/frontware/pwd)](https://pkg.go.dev/github.com/frontware/pwd)

-----------------------------------------------
<sup>© 2020 Frontware International. All Rights Reserved.</sup>