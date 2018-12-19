package viewer

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Config - config
type Config struct {
	BindAddr      string
	ServAddr      string
	JarvisNodeCfg string

	AnkaDB struct {
		DBPath string
		Engine string
	}
}

// LoadConfig - load config
func LoadConfig(filename string) (*Config, error) {
	fi, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = yaml.Unmarshal(fd, cfg)
	if err != nil {
		return nil, err
	}

	err = checkConfig(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// checkConfig - check config file
func checkConfig(cfg *Config) error {
	return nil
}
