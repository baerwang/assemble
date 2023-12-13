package config

import (
	"assemble/logger"
	_ "embed"
	"os"

	"gopkg.in/yaml.v3"
)

//go:embed config.yaml
var configContent []byte

var conf *Config

func ConfigInit() {

	var config Config

	if err := yaml.Unmarshal(configContent, &config); err != nil {
		panic("read config file fail : " + err.Error())
	}

	if file, err := os.ReadFile("config.yaml"); err == nil {
		if err = yaml.Unmarshal(file, &config); err != nil {
			panic("read config file fail : " + err.Error())
		}
	}

	conf = &Config{}
	conf.SetConfig(config)

	if conf == nil {
		panic("init config fail")
	}

	logger.InitLogger(true, config.Project.Level, "")
}

type Config struct {
	Project  Project
	Database DbConfig
	Redis    RedisConfig
	Minio    MinioConfig
	Applets  AppletsConfig
}

type Project struct {
	Name  string `yaml:"name"`
	Level int    `yaml:"level"`
}

type DbConfig struct {
	Host          string `yaml:"host"`
	Port          int    `yaml:"port"`
	DbName        string `yaml:"dbname"`
	UserName      string `yaml:"uname"`
	Password      string `yaml:"password"`
	OpenCount     int    `yaml:"open_count"`
	IdleCount     int    `yaml:"idle_count"`
	LifeTime      int    `yaml:"life_time"`
	IdleTime      int    `yaml:"idle_time"`
	SlowThreshold int    `yaml:"slow_threshold"`
}

type RedisConfig struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	DbName      int    `yaml:"dbname"`
	Password    string `yaml:"password"`
	MaxIdle     int    `yaml:"max_idle"`
	IdleTimeOut int    `yaml:"idle_time_out"`
}

type MinioConfig struct {
	Port            int    `yaml:"port"`
	Endpoint        string `yaml:"endpoint"`
	AccessKeyId     string `yaml:"access_key_id"`
	SecretAccessKey string `yaml:"secret_access_key"`
}

type AppletsConfig struct {
	AppId  string `yaml:"appId"`
	Secret string `yaml:"secret"`
}

func (r *Config) SetConfig(config Config) {
	r.Project = config.Project
	r.Database = config.Database
	r.Redis = config.Redis
	r.Minio = config.Minio
}

func GetConfig() *Config {
	return conf
}

func GetProject() Project {
	return conf.Project
}

func GetMinio() MinioConfig {
	return conf.Minio
}

func GetRedis() RedisConfig {
	return conf.Redis
}

func GetApplets() AppletsConfig {
	return conf.Applets
}
