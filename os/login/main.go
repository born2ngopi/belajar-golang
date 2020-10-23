package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Println(user)

	fmt.Println("real user: " + os.Getenv("SUDO_USER"))
}
