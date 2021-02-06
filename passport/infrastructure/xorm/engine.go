package xorm

import (
	"github.com/changqing98/cqzone/user/infrastructure/config"
    "github.com/go-xorm/xorm"
)
import _ "github.com/go-sql-driver/mysql"

var configInst = config.GetConfig()
var engineInst, _ = xorm.NewEngine(configInst.Db.Driver, configInst.Db.DataSource)

func GetXormEngine() *xorm.Engine {
    return engineInst
}
