package main

import (
	"fmt"
)

type User struct{
	userID, lastLogin int64
	firstName, secondName string
}

func main(){
	var Perl = User{1, 1633536132, "Roman", "Terekh"}
	var Yuki = User{2, 1633536056, "Yulia", "Terekh"}
	var Big = User{3, 1633536458, "Andrew", "Energy"}

	var Users = []User{Perl, Yuki, Big}

	for i, j := range Users{
		fmt.Println("Id", i, "|", j)
	}
}