package conf

import (
	"awesomeProject/model"
	"github.com/joho/godotenv"
	"os"
)

/**初始化数据库连接
 */

func Init() {
	godotenv.Load()
	//用来读取.env配置文件中MYSQL_DSN所对应的参数，即连接数据库的参数
	model.ConnectDB(os.Getenv("MYSQL_DSN"))
}
