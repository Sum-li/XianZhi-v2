package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:llc123..@tcp(106.14.146.41:3306)/ynk_cms?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	//设置连接的最大连接周期，超时自动关闭
	dbConn.SetConnMaxLifetime(100 * time.Second)
	//设置最大连接数
	dbConn.SetMaxOpenConns(100)
	//设置闲置连接数
	dbConn.SetMaxIdleConns(15)
}
