package schreibvogel

import (
	"context"

	"github.com/mattn/go-mastodon"
	log "github.com/sirupsen/logrus"
)

// Post stores posse data for a blog post.
type Post struct {
	URL   string          `json:"url"`
	POSSE map[string]Copy `json:"posse"`
}

// Copy stores metadata for a syndicated copy of a blog post.
type Copy struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type POSSEr interface {
	Post(text string) (Copy, error)
	Store(path string) error
}

type mast struct {
	server       string
	clientID     string
	clientSecret string
	username     string
	password     string
}

func (m *mast) Post(text string) (Copy, error) {
	c := mastodon.NewClient(&mastodon.Config{
		Server:       m.server,
		ClientID:     m.clientID,
		ClientSecret: m.clientSecret,
	})
	err := c.Authenticate(context.Background(), m.username, m.password)

	if err != nil {
		log.Error(err)
		return Copy{}, err
	}

	s, err := c.PostStatus(context.Background(), &mastodon.Toot{Status: text})
	if err != nil {
		log.Error(err)
		return Copy{}, err
	}

	return Copy{ID: string(s.ID), URL: s.URL}, nil
}

func (m *mast) Store(path string) error {
	return nil
}

func NewMast(server, clientID, clientSecret, username, password string) POSSEr {
	p := &mast{
		server:       server,
		clientID:     clientID,
		clientSecret: clientSecret,
		username:     username,
		password:     password,
	}

	return p
}

func Syndicate(config string) {
	// needEnvs := []string{"a"}
	// envs := os.Environ()
	log.Info("running syndication...")
}
