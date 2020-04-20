package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id   int
	name string
}

func main() {
	sqlConnStr := "root:root@tcp(127.0.0.1:3306)/db1"
	mysqlDB, err := sql.Open("mysql", sqlConnStr)
	if err != nil {
		fmt.Sprintf("数据库连接失败:%s\n", err)
		return
	}
	err = mysqlDB.Ping()
	if err != nil {
		fmt.Sprintf("数据库连接失败:%s\n", err)
		return
	}
	//queryUser := "select * from user"
	//rows, err := mysqlDB.Query(queryUser)
	//if err != nil {
	//	fmt.Sprintf("查询用户失败:%s\n", err)
	//	return
	//}
	//defer rows.Close()
	//for rows.Next() {
	//	var u = user{}
	//	err := rows.Scan(&u.id, &u.name)
	//	if err != nil {
	//		fmt.Sprintf("查询用户失败:%s\n", err)
	//		return
	//	}
	//	fmt.Println(u)
	//}

	//insertUserStr := " insert into user(id,name) value(?,?)"
	//deleteUserStr := "delete from user where id=?"
	//updateUserStr := "update user set name='测试' where id =?"
	//ret, err := mysqlDB.Exec(updateUserStr, 3)
	//if err != nil {
	//	fmt.Sprintf("插入用户失败:%s\n", err)
	//	return
	//}
	//fmt.Println(ret.RowsAffected())
	txDemo(mysqlDB)

}

// 事务
func txDemo(db *sql.DB){
	tx,err:=db.Begin()
	if err !=nil{
		if tx !=nil{
			tx.Rollback()
		}
		fmt.Sprintf("事务开启失败:%s/n",err)
		return
	}
	updateUserStr := "update user set name='测试' where id =?"
	ret, err := db.Exec(updateUserStr, 2)
	if err != nil {
		tx.Rollback()
		fmt.Sprintf("插入用户失败:%s\n", err)
		return
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Sprintf("插入用户失败:%s\n", err)
		return
	}
	fmt.Println(ret.RowsAffected())

}
