package conf

import (
	"log/slog"
	"sync"

	"github.com/chhz0/asyncgo/pkg/config"
)

var (
	conf *Config
	once sync.Once
)

type Config struct {
	Mode    string   `mapstructure:"mode"`
	TaskSvr *TaskSvr `mapstructure:"tasksvr"`
	MySQL   *MySQL   `mapstructure:"MySQL"`
	Redis   *Redis   `mapstructure:"Redis"`
}

type TaskSvr struct {
	Port string `mapstructure:"port"`
}

type MySQL struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"pass"`
	Db       string `mapstructure:"db_name" `
}

type Redis struct {
	Url         string `mapstructure:"url"`
	Auth        string `mapstructure:"auth"`
	MaxIdle     int    `mapstructure:"max_idle"`
	MaxActive   int    `mapstructure:"max_active"`
	IdleTimeout int    `mapstructure:"idle_timeout"`

	CacheTimeoutDay int `mapstructure:"cache_timeout_day"`
}

func GetConf() *Config {
	once.Do(initConfig)
	return conf
}

func initConfig() {
	config.New(
		config.WithConfigFile("tasksvr"),
		config.WithConfigFileType("toml"),
		config.WithConfigFilePath("../config"),
		config.WithUnmarshalStruct(&conf),
	).LoadConfig()

	slog.Info("init Config", slog.Group("Config", &conf))
}
