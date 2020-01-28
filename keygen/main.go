package keygen

import (
	"fmt"
	"go-practice/password"
)

func Execute()  {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Type a password to encrypt: ")
	// text, _ := reader.ReadString('\n')
	// hash, err := password.New(text)
	pwd := password.GetPwd(true)
	hash := password.HashAndSalt(pwd)
	// if (err != nil){
	// 	fmt.Println(err)
	// }
	// fmt.Printf("Password is %v \n encrypted as: %v ",text, hash )
	fmt.Printf("Password encrypted as: %v \n", hash)

	password.TestPwd([]byte(hash))
}