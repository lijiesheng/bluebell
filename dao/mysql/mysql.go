package mysql

// 参考的是 sqlx  https://www.liwenzhou.com/posts/Go/sqlx/
import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initMysql() (err error) {
	dsn := "root:ljs024816@tcp(127.0.0.1:3306)/bluebell?charset=utf8mb4&parseTime=True"
	db, err = sqlx.Connect("mysql", dsn) // 这个函数自带了 ping
	if err != nil {
		fmt.Println("connect DB failed", err)
		return
	}
	db.SetMaxOpenConns(20) // todo 暂时就设定这么多，有需要调整
	db.SetMaxIdleConns(10) // todo 暂时就设定这么多，有需要调整
	return
}
