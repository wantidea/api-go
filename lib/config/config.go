package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"time"
)

type app struct {
	RunMode         string
	Name            string
	Port            string
	JwtSecret       string
	JwtAdminSecret  string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	PageSize        int
	MaxPageSize     int64
	MD5Salt         string
	RuntimeRootPath string
	Locale          string
	MysqlState      bool
	RedisState      bool
	MongodbState    bool
	RabbitmqState   bool
	LoggerState     bool
}

type logger struct {
	SavePath   string
	SaveName   string
	FileExt    string
	TimeFormat string
}

type mysql struct {
	Connection string
	Host       string
	Port       string
	Database   string
	UserName   string
	Password   string
}

type redis struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type mongodb struct {
	Connection string
	Host       string
	Port       string
	Database   string
	UserName   string
	Password   string
}

type rabbitMQ struct {
	Connection string
	Host       string
	Port       string
	Database   string
	UserName   string
	Password   string
}

type oss struct {
}

type image struct {
	PrefixUrl string
	StaticUrl string
	SavePath  string
	MaxSize   int
	AllowExt  string
}

type video struct {
}

var AppConfig = &app{}
var LoggerConfig = &logger{}
var MysqlConfig = &mysql{}
var RedisConfig = &redis{}
var MongodbConfig = &mongodb{}
var RabbitMQConfig = &rabbitMQ{}
var OssConfig = &oss{}
var ImageConfig = &image{}
var VideoConfig = &video{}
var cfg *ini.File

func Setup() {
	configMap("app", AppConfig)
	configMap("image", ImageConfig)
	configMap("logger", LoggerConfig)
	configMap("mongodb", MongodbConfig)
	configMap("mysql", MysqlConfig)
	configMap("redis", RedisConfig)
	configMap("oss", OssConfig)
	configMap("rabbitmq", RabbitMQConfig)
	configMap("video", VideoConfig)
}

func configMap(path string, v interface{}) {
	var err error
	path = fmt.Sprintf("config/%s.ini", path)
	cfg, err = ini.Load(path)
	if err != nil {
		log.Fatalf("配置文件解析失败 '%s': %v", path, err)
	}

	err = cfg.Section("default").MapTo(v)
	if err != nil {
		log.Fatalf("配置块： %s ，遍历失败: %v", "default", err)
	}

	// 依据环境替换配置
	if AppConfig.RunMode != "" {
		err = cfg.Section(AppConfig.RunMode).MapTo(v)
		if err != nil {
			log.Fatalf("配置块： %s ，遍历失败: %v", AppConfig.RunMode, err)
		}
	}
}
