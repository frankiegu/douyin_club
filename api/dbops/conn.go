package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err error
)

func init() {
	//open的时候并不会去连接
	dbConn, err = sql.Open("mysql", "root:123456@tcp(39.107.77.94:3306)/video_server?charset=utf8")
	//dbConn.Ping() 测试DSN是否有问题
	if err != nil {
		panic(err.Error())
	}
}