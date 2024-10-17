package config

import (
	"fmt"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/loopholelabs/cmdutils/pkg/config"
)

var _ config.Config = (*Config)(nil)

var (
	configFile string
	logFile    string
)

const (
	defaultConfigFile = "flux.yml"
	defaultLogFile    = "flux.log"

	DefaultListenAddress = "127.0.0.1:8080"
	DefaultEndpoint      = "localhost:8080"
)

// Config is dynamically sourced from various files and environment variables.
type Config struct {
	ListenAddress string `mapstructure:"listen_address"`
	Endpoint      string `mapstructure:"endpoint"`
}

func New() *Config {
	return &Config{
		ListenAddress: DefaultListenAddress,
		Endpoint:      DefaultEndpoint,
	}
}

func (c *Config) RootPersistentFlags(_ *pflag.FlagSet) {}

func (c *Config) GlobalRequiredFlags(_ *cobra.Command) error {
	return nil
}

func (c *Config) Validate() error {
	err := viper.Unmarshal(c)
	if err != nil {
		return fmt.Errorf("unable to unmarshal config: %w", err)
	}
	return nil
}

func (c *Config) DefaultConfigDir() (string, error) {
	return xdg.ConfigHome, nil
}

func (c *Config) DefaultConfigFile() string {
	return defaultConfigFile
}

func (c *Config) DefaultLogDir() (string, error) {
	return xdg.StateHome, nil
}

func (c *Config) DefaultLogFile() string {
	return defaultLogFile
}

func (c *Config) SetConfigFile(file string) {
	configFile = file
}

func (c *Config) GetConfigFile() string {
	return configFile
}

func (c *Config) SetLogFile(file string) {
	logFile = file
}

func (c *Config) GetLogFile() string {
	return logFile
}
