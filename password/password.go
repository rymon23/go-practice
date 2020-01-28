// Package password provides functions for handling passwords.
package password

import (
	"fmt"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// cost is the bcrypt cost used when generating a new password.
const cost = bcrypt.DefaultCost

// New returns a new bcrypt password for the supplied password string.
func New(pass string) ([]byte, error) {
	if pass == "" {
		return nil, errors.New("password required")
	}

	p, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Authorize compares the password with the supplied crypt. If the
// result is nil, the account has been authorized.
func Authorize(crypt []byte, pass string) error {
	return bcrypt.CompareHashAndPassword(crypt, []byte(pass))
}

func TestPwd(crypt []byte) {
	// Prompt the user to enter a password
    fmt.Printf("Enter your new password to test: \n\n")
    // Variable to store the users input
    var pwd string
    // Read the users input
    _, err := fmt.Scan(&pwd)
    if err != nil {
        log.Println(err)
	}
	authErr := Authorize(crypt, pwd)

	if authErr != nil {
		fmt.Printf("Invalid password: '%v' \n", pwd)
		fmt.Println(err)
		return
	}
	fmt.Printf("Password '%v' validated! \n", pwd)
}

func GetPwd(displayPwd bool) []byte {
    // Prompt the user to enter a password
    fmt.Printf("Enter a password: \n\n")

    // Variable to store the users input
    var pwd string
    // Read the users input
    _, err := fmt.Scan(&pwd)
    if err != nil {
        log.Println(err)
	}
	if displayPwd {
		fmt.Printf("\n Password: '%v' \n", pwd)
	}
    // Return the users input as a byte slice which will save us
    // from having to do this conversion later on
    return []byte(pwd)
}

func HashAndSalt(pwd []byte) string {
    
    // Use GenerateFromPassword to hash & salt pwd.
    // MinCost is just an integer constant provided by the bcrypt
    // package along with DefaultCost & MaxCost. 
    // The cost can be any value you want provided it isn't lower
    // than the MinCost (4)
    hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
    if err != nil {
        log.Println(err)
    }
    // GenerateFromPassword returns a byte slice so we need to
    // convert the bytes to a string and return it
    return string(hash)
}