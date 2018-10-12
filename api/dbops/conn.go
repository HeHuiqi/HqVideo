package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
//包级变量重用数据库连接
var (
	dbConn *sql.DB
	err error
)

func init()  {
	//[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	dbSource := "root:h12345678@tcp(127.0.0.1:3306)/hq_video?charset-utf-8"
	dbConn ,err = sql.Open("mysql",dbSource)
	if err != nil {
		panic(err.Error())
	}
}
