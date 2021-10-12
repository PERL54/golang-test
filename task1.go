package main

import (
	"fmt"
)

type UserText struct{
	header, text string
	timeStamp int64
}

type User struct{
	userID, lastLogin int64
	firstName, secondName string
	userText UserText
}

func swapNames(u User) User {
	firstName:= u.firstName
	u.firstName = u.secondName
	u.secondName = firstName
	return u
}

func (u *User) setName(name string){
	u.firstName = name
}

func main(){
	var Perl = User{1, 1633536132, "Roman", "Terekh", UserText{"Heder!", "SampleText", 1633536132}}

	// Метод 
	fmt.Println(swapNames(Perl))

	// Функция
	Perl.setName("Nikita")
	fmt.Println(Perl)
}