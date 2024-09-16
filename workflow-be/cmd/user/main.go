package main

import (
	"fmt"
	"workflow/internal/user"
)

func main() {
	password, err := user.HashPassword("secret")
	if err != nil {
		fmt.Println("err=", err.Error())
	}
	fmt.Println("password=", password)
}