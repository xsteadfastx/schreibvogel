package schreibvogel

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	assert := assert.New(t)

	os.Setenv("SCHREIBVOGEL_SYNDICATION_MASTODON_PASSWORD", "supersecret")
	defer os.Unsetenv("SCHREIBVOGEL_SYNDICATION_MASTODON_PASSWORD")

	tables := []struct {
		file     string
		expected map[string]map[string]string
	}{
		{
			file: "testdata/config-1.toml",
			expected: map[string]map[string]string{
				"mastodon": {
					"username": "marvin@xsteadfastx.org",
					"password": "supersecret",
				},
			},
		},
	}

	for _, table := range tables {
		c := &Config{}
		err := c.Parse(table.file)
		assert.Nil(err)
		log.Printf("%+v", c)
		assert.Equal(&Config{Syndication: table.expected}, c)
	}
}

func TestFromEnv(t *testing.T) {
	assert := assert.New(t)

	os.Setenv("ZONK_FOO_BAR_BLA", "yeah")

	defer os.Unsetenv("ZONK_FOO_BAR_BLA")

	r, err := FromEnv("zonk", "foo", "bar", "bla")
	assert.Nil(err)
	assert.Equal("yeah", r)
}

func TestFromEnvNoEnvVariable(t *testing.T) {
	assert := assert.New(t)
	r, err := FromEnv("zonk", "foo", "bar", "bla")
	assert.NotNil(err)
	assert.Equal("", r)
}
