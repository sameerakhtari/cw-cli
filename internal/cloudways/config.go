package cloudways

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config structure for saving credentials
type Config struct {
	Email       string `mapstructure:"email"`
	APIKey      string `mapstructure:"api_key"`
	AccessToken string `mapstructure:"access_token"`
}

// SaveConfig saves the credentials to a YAML config file
func SaveConfig(email, apiKey, token string) error {
	config := Config{
		Email:       email,
		APIKey:      apiKey,
		AccessToken: token,
	}

	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	cwcliDir := filepath.Join(configDir, "cw-cli")
	err = os.MkdirAll(cwcliDir, os.ModePerm)
	if err != nil {
		return err
	}

	configFile := filepath.Join(cwcliDir, "config.yaml")

	viper.Set("email", config.Email)
	viper.Set("api_key", config.APIKey)
	viper.Set("access_token", config.AccessToken)

	err = viper.WriteConfigAs(configFile)
	if err != nil {
		return err
	}

	fmt.Println("✅ Config saved to", configFile)
	return nil
}

// LoadConfig loads the credentials from config.yaml
func LoadConfig() (*Config, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	configFile := filepath.Join(configDir, "cw-cli", "config.yaml")
	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
