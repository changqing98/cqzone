package xorm

import "github.com/go-xorm/xorm"
import _ "github.com/go-sql-driver/mysql"

const DriverNameMysql = "mysql"

const DataSourceName = "changqing:123456@tcp(mysql:3306)/alien"

func GetXORMEngine() *xorm.Engine {
    engine, err := xorm.NewEngine(DriverNameMysql, DataSourceName)
    if err != nil {
        panic(err)
    }
    return engine
}
