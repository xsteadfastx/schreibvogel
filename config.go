package schreibvogel

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Syndication map[string]map[string]string `toml:"syndication"`
}

func (c *Config) Parse(file string) error {
	logger := log.WithFields(log.Fields{"config": file})
	_, err := toml.DecodeFile(file, c)

	if err != nil {
		logger.Fatal(err)
	}

	for service, conf := range c.Syndication {
		for k, v := range conf {
			if v == "" {
				logger.WithFields(
					log.Fields{"service": service, "key": k}).Info(
					"couldnt find value... looking in environment variables.",
				)

				lv, err := FromEnv(EnvPrefix, "syndication", service, k)

				if err != nil {
					return err
				}

				c.Syndication[service][k] = lv
			}
		}
	}

	return nil
}

// TODO: Make it Variadic.
func FromEnv(prefix, section, subsection, key string) (string, error) {
	env := fmt.Sprintf(
		"%s_%s_%s_%s",
		strings.ToUpper(prefix),
		strings.ToUpper(section),
		strings.ToUpper(subsection),
		strings.ToUpper(key),
	)

	log.WithFields(log.Fields{"var": env}).Info("looking for env variable")

	v := os.Getenv(env)

	if v == "" {
		return "", fmt.Errorf("could not find environment variable %q", env)
	}

	return v, nil
}
