package ciscoasa

import "github.com/CiscoDevNet/go-ciscoasa/ciscoasa"

// Config is the configuration structure used to instantiate a
// new CiscoAsa client.
type Config struct {
	APIURL      string
	Username    string
	Password    string
	SSLNoVerify bool
}

// NewClient returns a new CiscoASA client.
func (c *Config) NewClient() (*ciscoasa.Client, error) {
	return ciscoasa.NewClient(c.APIURL, c.Username, c.Password, c.SSLNoVerify)
}
