package config

import (
    "github.com/changqing98/cqzone/user/infrastructure/utils/logger"
    "github.com/ghodss/yaml"
    "io/ioutil"
    "os"
    "path/filepath"
    "sync"
    "time"
)

type userServiceConfig struct {
    DbConfig dbProperties `yaml:"db"`
}

type dbProperties struct {
    Host           string
    Port           string
    Username       string
    Password       string
    Database       string
    MaxLifeTime time.Duration

    MaxConnections int
}

var serviceConfig *userServiceConfig

var mutex sync.Mutex

func GetUserServiceConfig() *userServiceConfig {
    if serviceConfig == nil {
        mutex.Lock()
        if serviceConfig == nil {
            path, _ := filepath.Abs(os.Args[0])
            var yamlFilePath = filepath.Join(filepath.Dir(path), "/resource/application.yaml")
            logger.Info("yaml config file path: " + yamlFilePath)
            yamlFile, _ := ioutil.ReadFile(yamlFilePath)
            serviceConfig = &userServiceConfig{}
            yaml.Unmarshal(yamlFile, serviceConfig)
        }
    }
    return serviceConfig
}
