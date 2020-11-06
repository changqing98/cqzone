package config

import "database/sql"

const MYSQL = "mysql"

var dbConfig = GetUserServiceConfig().DbConfig

func getDataSource() *sql.DB {
    username := dbConfig.Username
    password := dbConfig.Password
    host := dbConfig.Host
    port := dbConfig.Port
    database := dbConfig.Database
    maxConnections := dbConfig.MaxConnections
    var dataSourceName = username + ":" + password + "@" + host + ":" + port + "/" + database
    db, err := sql.Open(MYSQL, dataSourceName)
    if err != nil {
        panic(err)
    }
    db.SetMaxOpenConns(maxConnections)
    return db
}
