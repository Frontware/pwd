// Copyright Frontware International
// This package is used by several Frontware project to handle basic tasks about passwords

package pwd

import (
	"bytes"
	"crypto/rand"
	_ "embed"
	"hash/fnv"
	"io"

	"github.com/steakknife/bloomfilter"
	"golang.org/x/crypto/bcrypt"
)

//go:embed pwd.bf.gz
var pwdBin []byte

var bf *bloomfilter.Filter

func init() {
	r := bytes.NewReader(pwdBin)
	bf, _, _ = bloomfilter.ReadFrom(r) // read the BF
}

// IsCommon returns true if the password is in the list of most common passwords.
// We use bloomfilter with a list of 1 million common passwords.
//  IsCommon("qwerty") // returns true
//  IsCommon("password") // returns true
//  IsCommon("jkljfd5675fhgf6567H=") // returns false
func IsCommon(pwd string) bool {
	h := fnv.New64()
	h.Write([]byte(pwd))

	return bf.Contains(h)
}

// HashPassword create a bcrypt hash
func HashPassword(pwd string) (string, error) {
	password := []byte(pwd)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// stdChars characters to be used in auto generated password
// We removed 1 and I to avoid confusion, sometimes they are very similar on the phone
var stdChars = []byte("ABCDEFGHJKLMNPQRSTUVWXYZabcdefghjkmnopqrstuvwxyz23456789")

// NewPassword returns a new password. Minimum length is 5
//  NewPassword(4) // returns a random password of 5 chars
//  NewPassword(8) // returns a random password of 8 chars
func NewPassword(length int) string {
	if length < 5 {
		length = 5
	}
	return randChar(length, stdChars)
}

// randChar returns random char, usually to generate password
func randChar(length int, chars []byte) string {
	newPword := make([]byte, length)
	randomData := make([]byte, length+(length/4)) // storage for random bytes.
	clen := byte(len(chars))
	maxrb := byte(256 - (256 % len(chars)))
	i := 0
	for {
		io.ReadFull(rand.Reader, randomData)
		for _, c := range randomData {
			if c >= maxrb {
				continue
			}
			newPword[i] = chars[c%clen]
			i++
			if i == length {
				return string(newPword)
			}
		}
	}
}
