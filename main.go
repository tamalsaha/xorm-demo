package main

import (
	"fmt"
	"github.com/masudur-rahman/xorm-demo/go-xorm"
	"github.com/masudur-rahman/xorm-demo/xormio"
)

func main() {
	fmt.Println("Running using `github.com/go-xorm/xorm`\n ")
	go_xormTest()

	println("\n\n\n\n")

	fmt.Println("Running using `xorm.io/xorm`\n ")
	xormioTest()
}

func go_xormTest() {
	go_xorm.Connect()
	defer go_xorm.Close()

	user, err := go_xorm.CreateUser(&go_xorm.User{
		Name:  "masud",
		Email: "masud@appscode.com",
		Address: go_xorm.Address{
			City:    "Dhaka",
			Country: "Bangladesh",
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", user)

	users, err := go_xorm.ListUser()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", users)
}

func xormioTest() {
	xormio.Connect()
	defer xormio.Close()

	user, err := xormio.CreateUser(&xormio.User{
		Name:  "masud",
		Email: "masud@appscode.com",
		Address: xormio.Address{
			City:    "Dhaka",
			Country: "Bangladesh",
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", user)

	users, err := xormio.ListUser()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", users)
}
