package pkg

import (
	"net/url"
)

func IsValidURL(rawURL string) bool {
	u, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return false
	}

	// Garante que a URL tem um esquema (http/https) e um host
	return (u.Scheme == "http" || u.Scheme == "https") && u.Host != ""
}
