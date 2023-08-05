package main

import (
	"fmt"
	"game-app/entity"
	"game-app/repository/mysql"
)

func main() {

	TestUserMysqlRepo()

}
func TestUserMysqlRepo() {
	mysqlRepo := mysql.New()

	createdUser, err := mysqlRepo.RegisterUser(entity.User{
		ID:          0,
		PhoneNumber: "091988855",
		Name:        "Athena Helali",
	})

	if err != nil {
		fmt.Println("register user", err)
	} else {
		fmt.Println("created user:", createdUser)
	}

	isUnique, err := mysqlRepo.IsPhoneNumberUnique(createdUser.PhoneNumber + "2")
	if err != nil {
		fmt.Println("unique error", err)
	}

	fmt.Println("is unique", isUnique)

}
