package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	_, err := sql.Open("mysql", "root:@/sckshuhari/tcp(127.0.0.1:3306)/testsck")
	if err != nil {
		fmt.Println("connect fail")

	}
	fmt.Println("connect success")
}
