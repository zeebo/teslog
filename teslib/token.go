package teslib

import "time"

type Token struct {
	AccessToken  string
	RefreshToken string
	Expires      time.Time
}

func (c *Token) ShouldRefresh() bool {
	return time.Now().Add(24 * time.Hour).After(c.Expires)
}

func (c *Token) Expired() bool {
	return time.Now().After(c.Expires)
}
