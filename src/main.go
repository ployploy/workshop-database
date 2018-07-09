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
	//	add(db)
	//	fmt.Println(read(db))
	remove(db, "4")
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

func add(db *sql.DB) bool {
	results, _ := db.Prepare(`Insert INTO user
	(citizen_id,firstname,lastname,birthyear
		,firstname_father,lastname_father,firstname_mother,lastname_mother,soldier_id
		,address_id)
		VALUES(?,?,?,?,?,?,?,?,?,?)`) //เป็นการเตรียมค่าเพื่อจะทำงาน

	_, err := results.Exec("1092018019065", "พอยพอย", "ซันชาย", 1995, "ตะวัน", "ซันชาย", "สมิตา", "ซันชาย", 988, 1)
	if err != nil {
		panic(err)
		return false
	}
	return true
}

func remove(db *sql.DB, id string) bool {
	statement, _ := db.Prepare("DELETE FROM testsck.user WHERE user_id=?")

	_, err := statement.Exec(id)
	if err != nil {
		panic(err.Error())
		return false
	}
	return true
}
