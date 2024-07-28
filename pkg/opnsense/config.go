package opnsense

import "net/url"

type Authentication struct {
	Username string
	Password string
}

type Config struct {
	Authentication Authentication
	BaseURL string
}

type option func(Config) (Config, error)

func NewConfig(o ...option) (Config, error) {
	c := Config{}
	var err error

	for _, opt := range o {
		c, err = opt(c)
		if err != nil {
			return Config{}, err
		}
	}
	return c, nil
}

func WithURL(u string) option {
	return func(c Config) (Config, error) {
		newURL, err := url.Parse(u)
		if err != nil {
			return Config{}, err
		}
		c.BaseURL = newURL.Scheme + "://" + newURL.Host

		return c, nil
	}
}

func WithAuthentication(u string, p string) option {
	return func(c Config) (Config, error) {
		c.Authentication.Username = u
		c.Authentication.Password = p
		return c, nil
	}
}
