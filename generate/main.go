package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/frontware/pwd"
	"github.com/steakknife/bloomfilter"
)

const (
	bmFilePath = "../pwd.bf.gz"
	txtFile    = "/tmp/pwd.txt"

	maxElements = 1000000
	probCollide = 0.01
)

// GenerateBloom generates bloom file based on pwd.txt imported from 10 millions top pwd list.
func GenerateBloom() {

	list := "https://raw.githubusercontent.com/danielmiessler/SecLists/master/Passwords/Common-Credentials/10-million-password-list-top-1000000.txt"

	DownloadFile(txtFile, list)

	file, err := os.Open(txtFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bf, err := bloomfilter.NewOptimal(maxElements, probCollide)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	// Scan every line of the file
	for scanner.Scan() {
		h := fnv.New64()
		h.Write(scanner.Bytes())
		// Add hash64 to bloom filter
		bf.Add(h)
	}
	_, err = bf.WriteFile(bmFilePath)
	fmt.Printf("File %s created\n", bmFilePath)
}

func main() {

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 0 {
		fmt.Println("Tool command line:")
		fmt.Println()
		fmt.Println("- generate : Generate new bloom filter file")
		fmt.Println("- try password : Check if password is common or not")
		fmt.Println()
		os.Exit(0)
	}

	switch argsWithoutProg[0] {

	case "generate":

		GenerateBloom()

	case "try":
		if len(argsWithoutProg) < 2 {
			fmt.Println("You must specify the password")
			os.Exit(0)
		}
		fmt.Printf("Password %s common : %v\n", argsWithoutProg[1], pwd.IsCommon(argsWithoutProg[1]))
	}

}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
