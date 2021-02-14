package config

type config struct {
    Db struct {
        Driver     string
        DataSource string
    }
}

var configInst = &config{
    Db: struct {
        Driver     string
        DataSource string
    }{Driver: "mysql", DataSource: "changqing:123456@tcp(mysql:3306)/passport"},
}

func GetConfig() *config {
    return configInst
}
