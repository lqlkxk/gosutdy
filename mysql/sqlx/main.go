package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)
var mysqlDb *sqlx.DB
// sqlx可访问
type User struct {
	Id   int `db:"id"`
	Name string `db:"name"`
}

func main() {
	var users []*User
	initDb()
	var user User
	err := mysqlDb.Get(&user, "select id, name from user where id=?", 2)
	if err == sql.ErrNoRows {
		log.Printf("not found data of the id:%d", 1)
	}

	if err != nil {
		panic(err)
	}

	fmt.Printf("user: %#v\n", user)
	fmt.Println(user)
	err = mysqlDb.Select(&users, "select id, name from user")
	if err != nil {
		fmt.Sprintf("查询用户失败:%s/n", err)
	}

	for _,v := range users{
		log.Println(v.Name)
	}
}
func initDb(){
	sqlConnStr := "root:root@tcp(127.0.0.1:3306)/db1"
	db, err := sqlx.Connect("mysql", sqlConnStr)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	mysqlDb = db
	return
}
