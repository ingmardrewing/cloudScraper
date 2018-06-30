package cloudScraper

import (
	"log"
	"regexp"
)

func NewCloudScraper(data string) CloudScraper {
	cs := new(cloudScraper)
	cs.data = data
	return cs
}

type CloudScraper interface {
	GetFirstCapturingGroupValue() string
	SetPattern(string)
}

type cloudScraper struct {
	data    string
	pattern *regexp.Regexp
}

// Sets the regex pattern to use for
// the data extraction.
// The pattern must contain a capturing group
func (c *cloudScraper) SetPattern(pattern string) {
	if len(pattern) == 0 {
		log.Fatalln("Pattern must not be an empty string")
	}
	c.pattern = regexp.MustCompile(pattern)
}

// Returns the first capturing group
// defined within the pattern, if any.
// Retuns an empty string if nothing is found
func (c *cloudScraper) GetFirstCapturingGroupValue() string {
	matches := c.match()
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// compiles the given pattern and returns
// the matches using FindStringSubmatch
func (c *cloudScraper) match() []string {
	return c.pattern.FindStringSubmatch(c.data)
}
