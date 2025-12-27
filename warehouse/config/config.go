package config

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v4"
)

type Config struct {
	Database                 DatabaseConfig           `yaml:"database"`
	Kafka                    KafkaConfig              `yaml:"kafka"`
	WarehouseServiceSettings WarehouseServiceSettings `yaml:"WarehouseServiceSettings"`
	Redis                    RedisConfig              `yaml:"redis"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"name"`
	SSLMode  string `yaml:"ssl_mode"`
}

type KafkaConfig struct {
	Host                           string `yaml:"host"`
	Port                           int    `yaml:"port"`
	AskForGGOrderInfoTopicName     string `yaml:"ask_for_ggorder_info"`
	AskForGGOrderInfoEditTopicName string `yaml:"ask_for_ggorder_info_edit"`
	ProvideGGOrderInfoTopicName    string `yaml:"provide_ggorder_info"`
}

type WarehouseServiceSettings struct {
	MinNameLen int `yaml:"minNameLen"`
	MaxNameLen int `yaml:"maxNameLen"`
}

type RedisConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	DB   int    `yaml:"db"`
	TTL  int    `yaml:"ttlSeconds"`
}


func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	return &config, nil
}

func (c Config) RedisAddr() string {
	return fmt.Sprintf("%s:%d", c.Redis.Host, c.Redis.Port)
}

