package mysql

// 参考的是 sqlx  https://www.liwenzhou.com/posts/Go/sqlx/
import (
	"bluebell/setting"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitMysql(cfg *setting.Mysql) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname)
	DB, err = sqlx.Connect("mysql", dsn) // 这个函数自带了 ping
	if err != nil {
		fmt.Println("connect DB failed", err)
		return
	}
	DB.SetMaxOpenConns(cfg.MaxOpenConns) // todo 暂时就设定这么多，有需要调整
	DB.SetMaxIdleConns(cfg.MaxIdleConns) // todo 暂时就设定这么多，有需要调整
	return
}

func Close() {
	_ = DB.Close()
}
