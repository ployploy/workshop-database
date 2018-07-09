package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID              int
	CitizenID       string
	Firstname       string
	Lastname        string
	BirthYear       int
	FirstnameFather string
	LastnameFather  string
	FirstnameMother string
	LastnameMother  string
	SoldierID       int
	AddressID       int
}

func main() {
	db, err := sql.Open("mysql", "root:sckshuhari@tcp(127.0.0.1:3306)/testsck")
	defer db.Close()
	if err != nil {
		fmt.Println("connect fail")

	}
	fmt.Println("connect success")
	fmt.Println(read(db))

}

func read(db *sql.DB) []User {
	results, _ := db.Query("SELECT * FROM user")
	var userList []User
	for results.Next() {
		var user User
		err := results.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.CitizenID, &user.BirthYear, &user.FirstnameFather, &user.LastnameFather, &user.FirstnameMother, &user.LastnameMother, &user.SoldierID, &user.AddressID)
		if err != nil {
			panic(err)
		}
		userList = append(userList, user)
	}
	return userList
}
