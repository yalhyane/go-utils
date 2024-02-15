package url_utils

import (
	"net/url"
	"strings"
)

// GetUrlDomain Parse domain from url
func GetUrlDomain(u string) string {
	p, err := url.Parse(u)
	if err != nil {
		return ""
	}
	return strings.TrimPrefix(p.Hostname(), "www.")
}
