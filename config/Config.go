package config

import (
	"bufio"
	"encoding/json"
	"os"
)

/*
conf, err := config.ParseConfig("./config/app.json")
if err != nil {
	panic("读取配置文件失败"+ err.Eroor())
}

fmt.Printf("conf:%#v\n", config)

// 在接口中输出整个配置
type UserController struct{}

func (CONTROLLER *UserController) Get(context *gin.Context) {
	id := context.Query("id")
	context.JSON(http.StatusOK, gin.H{
		"id": id,
		"conf": config.GetConfig(),
	})
}
*/

type Config struct {
	AppName     string         `json:"app_name,omitempty"`
	AppModel    string         `json:"app_model,omitempty"`
	AppHost     string         `json:"app_host,omitempty"`
	AppPort     string         `json:"app_port,omitempty"`
	Database    DatabaseConfig `json:"database"`
	RedisConfig RedisConfig    `json:"redis_config,omitempty"`
}

type DatabaseConfig struct {
	Driver   string `json:"addr"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"db_name"`
	Charset string `json:"charset"`
	ShowSql  bool   `json:"show_sql"`
}

type RedisConfig struct {
	Addr     string `json:"addr"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

// 获取配置，外部使用 config.GetConfig() 调用
func GetConfig() *Config {
	return cfg
}
// 存储配置的全局对象
var cfg *Config = nil

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path) // 读取文件
	defer file.Close()
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader) // 解析json
	if err = decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
