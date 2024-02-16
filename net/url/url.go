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

// ResolveLink resolves a given link against a base link by parsing and checking if the link is absolute,
// and then using the `ResolveReference` method to resolve the relative link.
func ResolveLink(baseLink, link string) (string, error) {
	// Parse the base link
	baseParsed, err := url.Parse(baseLink)
	if err != nil {
		return "", err
	}

	// Parse the provided link
	linkParsed, err := url.Parse(link)
	if err != nil {
		return "", err
	}

	// If the link is already absolute, return it as is
	if linkParsed.IsAbs() {
		return link, nil
	}

	// Resolve the relative link against the base link
	resolvedURL := baseParsed.ResolveReference(linkParsed)

	return resolvedURL.String(), nil
}
