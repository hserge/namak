package app

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Db struct {
		Dsn string `yaml:"dsn" ,envconfig:"APP_DB_DSN"`
	} `yaml:"db"`
	Server struct {
		Port string `yaml:"port" ,envconfig:"APP_SERVER_PORT"`
	} `yaml:"server"`
}

type Configurable interface {
	Load() Config
	readFile(cfg *Config)
	readEnv(cfg *Config)
}

var configFile string = "config.yml"

func New(cfgFile ...string) (*Config, error) {
	if len(cfgFile) == 1 {
		configFile = cfgFile[0]
	}
	var cfg Config
	err := cfg.Load()
	return &cfg, err
}

func (c *Config) Load() error {
	// load yaml config first
	err := c.readFile(configFile)
	if err != nil {
		return err
	}

	// load env and overwrite yaml config
	return c.readEnv()
}

func (c *Config) readFile(cfgFile string) error {
	f, err := os.Open(cfgFile)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	return decoder.Decode(c)
}

func (c *Config) readEnv() error {
	return envconfig.Process("APP_", c)
}
