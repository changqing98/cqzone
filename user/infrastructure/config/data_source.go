package config

import "github.com/go-xorm/xorm"
import _ "github.com/go-sql-driver/mysql"

const DriverNameMysql = "mysql"

const DataScourceName = "changqing:123456@tcp(mysql:3306)/user"

func GetXORMEngine() *xorm.Engine {
    engine, err := xorm.NewEngine(DriverNameMysql, DataScourceName)
    if err != nil {
        panic(err)
    }
    return engine
}
