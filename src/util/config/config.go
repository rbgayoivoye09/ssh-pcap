package config

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"

	. "github.com/rbgayoivoye09/ssh-pcap/src/util/log"
)

type Config struct {
	LocalFilePath string                     `yaml:"local-file-path" mapstructure:"local-file-path"`
	SSHConfig     map[string]SSHServerConfig `yaml:"ssh-config" mapstructure:"ssh-config"`
}

type SSHServerConfig struct {
	Host     string `yaml:"host" mapstructure:"host"`
	Port     string `yaml:"port" mapstructure:"port"`
	Username string `yaml:"username" mapstructure:"username"`
	Password string `yaml:"password" mapstructure:"password"`
	PcapCmd  string `yaml:"pcapCmd" mapstructure:"pcapCmd"`
}

// newConfig creates a new instance of Config using the provided config file path.
func newConfig(configFilePath string) (*Config, error) {
	// Set the config file path for viper.
	viper.SetConfigFile(configFilePath)

	// Read the config file.
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	// Unmarshal the config file into a Config struct.
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %v", err)
	}

	// Return the config instance.
	return &config, nil
}

// GetConfig retrieves the configuration settings.
//
// It returns a pointer to a Config struct.
func GetConfig(cfgPath string) *Config {

	Logger.Sugar().Info("GetConfig: ", cfgPath)

	configFilePath := ""

	if cfgPath != "" {
		configFilePath = cfgPath
	} else {
		Logger.Sugar().Error("cfgPath is empty")
		configFilePath = getDefaultConfigFilePath()
	}

	// Read the config file and create a new Config object.
	config, err := newConfig(configFilePath)
	if err != nil {
		// Log and exit if there is an error reading the config file.
		Logger.Sugar().Fatalf("Error reading config file: %v", err)
	}

	return config
}

func getDefaultConfigFilePath() string {
	// Get the absolute path of the project root directory.
	projectRoot, err := filepath.Abs(".")
	if err != nil {
		// Log and exit if there is an error getting the project root path.
		Logger.Sugar().Fatalf("Error getting project root path: %v", err)
	} else {
		Logger.Sugar().Infof("project root: %s", projectRoot)
	}
	// Construct the path to the config file.
	configFilePath := filepath.Join(projectRoot, "config", "base.yaml")
	return configFilePath
}
