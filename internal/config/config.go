package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
	Database    `yaml:"database"`
	S3Server    `yaml:"s3server"`
	RedisServer `yaml:"redis"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"5s"`
	MaxSize     int64         `yaml:"max_size" env-default:"5"`
	UploadDir   string        `yaml:"upload_dir" env-default:"/tmp/image-converter"`
}

type Database struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"5432"`
	User     string `yaml:"user" env-default:"postgres"`
	Password string
	DBName   string `yaml:"db_name" env-default:"postgres"`
	SSLMode  string `yaml:"ssl_mode" env-default:"disable"`
}

type S3Server struct {
	Endpoint  string `yaml:"endpoint"`
	Region    string `yaml:"region"`
	Bucket    string `yaml:"bucket_name"`
	AccessKey string
	SecretKey string
}

type RedisServer struct {
	Address string `yaml:"address" env-default:"localhost:6379"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file %s does not exist", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	if _, err := os.Stat(cfg.UploadDir); os.IsNotExist(err) {
		if err = os.Mkdir(cfg.UploadDir, os.ModePerm); err != nil {
			log.Fatalf("cannot create directory for tmp images: %s", err)
		}
	}

	cfg.Database.Password = os.Getenv("DB_PASSWORD")
	cfg.S3Server.SecretKey = os.Getenv("S3_SECRET_KEY")
	cfg.S3Server.AccessKey = os.Getenv("S3_ACCESS_KEY")

	return &cfg
}
