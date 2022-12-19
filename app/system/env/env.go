package env

import (
    "sync"

    "github.com/jinzhu/configor"
)

type Config struct {
    CreditDBHost     string `required:"true" env:"Credit_DB_HOST"`
    CreditDBPort     string `required:"true" env:"Credit_DB_PORT"`
    CreditDBUser     string `requried:"true" env:"Credit_DB_USER"`
    CreditDBPassword string `required:"true" env:"Credit_DB_PASSWORD"`
}

var (
    configOnce sync.Once
    configInst *Config
)

func GetConfig() *Config {
    configOnce.Do(func() {
        if err := configor.Load(&configInst); err != nil {
            panic(err)
        }
    })
    return configInst
}