package sqlservice

import (
	"database/sql"
	"fmt"
	"log"

	//mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

//数据库用户名
const username = ""

//数据库密码
const password = ""

//数据库地址(端口号)
const host = "localhost"

//数据库名称
const dbname = "db_dnf"

//端口号
const port = "3306"

var connection *sql.DB

func init() {
	connect()
}

func connect() {
	var err1 error
	connection, err1 = sql.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+dbname)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println("数据库连接正常")
}

//GetSave 获取存档,成功返回存档字符串,失败则返回空串
func GetSave(uid string) string {
	row, err := connection.Query("select save from tb_save where userid=?", uid)
	if err != nil {
		log.Fatal(err)
	}
	savestr := ""
	for row.Next() {
		if err := row.Scan(&savestr); err != nil {
			log.Fatal(err)
		}
		break
	}
	return savestr
}

//SaveSave 保存存档
func SaveSave(uid string, savestr string) bool {
	var err error
	// var result sql.Result
	if GetSave(uid) == "" {
		_, err = connection.Exec("insert into tb_save value(?,?)", uid, savestr)
	} else {
		_, err = connection.Exec("update tb_save set save=? where userid=?", savestr, uid)
	}
	if err != nil {
		log.Fatal(err)
	}
	return true
}

//IsUser 判断用户是否正确,如果正确返回ID，否则返回空的字符串
func IsUser(username string, password string) string {
	row, err := connection.Query("select id from tb_user where username=? and password=?", username, password)
	if err != nil {
		log.Fatal(err)
	}
	id := ""
	for row.Next() {
		if err := row.Scan(&id); err != nil {
			log.Fatal(err)
		}
		break
	}
	return id
}
