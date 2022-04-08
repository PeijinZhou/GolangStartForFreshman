package mysqlconnection

//用于连接数据库
import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func DatebaseInit() *sqlx.DB {
	//orm.NewOrm().
	database, err := sqlx.Open("mysql", "root:root@tcp(xxx.xxx.x.xxx:xxxx)/zpjdemo")
	fmt.Println("Connect to database succeed")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return nil
	}
	return database
}

//此函数可以连接数据库,并将返回一个实例化的数据库变量.
