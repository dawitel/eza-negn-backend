package config

import (
    "log"
    "github.com/spf13/viper"
)

// ConfigManager is responsible for loading and providing configuration
type ConfigManager struct {
    Config *Config
}

type Config struct {
    Server   ServerConfig   `mapstructure:"server"`
    Database DatabaseConfig `mapstructure:"database"`
    Logger   LoggerConfig   `mapstructure:"logger"`
    Secrets  SecretsConfig  `mapstructure:"secrets"`
}

type ServerConfig struct {
    Host string `mapstructure:"host"`
    Port string `mapstructure:"port"`
    Mode string `mapstructure:"mode"`
}

type DatabaseConfig struct {
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    User     string `mapstructure:"user"`
    Password string `mapstructure:"password"`
    Name     string `mapstructure:"name"`
    SSLMode  string `mapstructure:"sslmode"`
}

type LoggerConfig struct {
    Level string `mapstructure:"level"`
}

type SecretsConfig struct {
    JWTSecret string `mapstructure:"jwt_secret"`
}

// NewConfigManager creates a new ConfigManager and loads the configuration for the given environment
func NewConfigManager(env string) *ConfigManager {
    viper.AddConfigPath("../../configs")
    viper.SetConfigName(env)
    viper.SetConfigType("yaml")
    
    // Environment variable overrides
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file: %v", err)
    }

    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        log.Fatalf("Unable to decode into struct: %v", err)
    }

    return &ConfigManager{Config: &config}
}

// GetConfig returns the loaded configuration
func (cm *ConfigManager) GetConfig() *Config {
    return cm.Config
}
