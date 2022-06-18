package main

import (
	"database/sql"
	"fmt"
	"strings"
)

func main() {
	var conf []string
	var db *sql.DB
	var err error
	conf = append(conf, "Provider=SQLOLEDB")
	conf = append(conf, "Data Source=127.0.0.1\\SQLEXPRESS") // sqlserver IP 和 服务器名称
	conf = append(conf, "Initial Catalog=studb")             // 数据库名
	conf = append(conf, "user id=testlogin")                 // 登陆用户名
	conf = append(conf, "password=abc")                      // 登陆密码
	fmt.Println(strings.Join(conf, ";"))
	db, err = sql.Open("adodb", strings.Join(conf, ";"))
	if err != nil {
		fmt.Println("sql open:", err.Error())
		return
	}
	// 执行SQL语句
	rows, err := db.Query("select * from S")
	if err != nil {
		fmt.Println("query: ", err)
		return
	}
	for rows.Next() {
		var id int64
		var name string
		var age int64
		rows.Scan(&id, &name, &age)
		fmt.Printf("Name: %s \t id: %d age: %d\n", name, id, age)
	}
}
