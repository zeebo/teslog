package teslib

import (
	"net/url"
	"strings"
)

func makeURL(path string) string {
	split := strings.Index(path, "?")
	if split == -1 {
		split = len(path)
		path += "?"
	}

	return (&url.URL{
		Scheme:   "https",
		Host:     "owner-api.teslamotors.com",
		Path:     path[:split],
		RawQuery: path[split+1:],
	}).String()
}
