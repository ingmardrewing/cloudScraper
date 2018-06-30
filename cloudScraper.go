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
	GetFirstValueOfGroupNamed(groupName string) string
	GetValuesOfGroupsNamed(groupName string) []string
	GetAllNamedGroupMaps() []map[string]string
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

// Returns the value of a named capturing group
// defined within the regex, like the name "year" in
// this pattern: 'yadda yadda (?P<year>\d{4}) yadda'
func (c *cloudScraper) GetFirstValueOfGroupNamed(groupName string) string {
	groupMap := c.GetAllNamedGroupMaps()[0]
	if match, ok := groupMap[groupName]; ok {
		return match
	}
	return ""
}

// Returns an array of all matches found for the
// capturing group, which matches the param groupName
func (c *cloudScraper) GetValuesOfGroupsNamed(groupName string) []string {
	values := []string{}
	for _, m := range c.GetAllNamedGroupMaps() {
		if match, ok := m[groupName]; ok {
			values = append(values, match)
		}
	}
	return values
}

// Returns an array of maps of all group names as keys
// and their matches as values found in all matches
func (c *cloudScraper) GetAllNamedGroupMaps() []map[string]string {
	names := c.pattern.SubexpNames()
	matches := c.pattern.FindAllStringSubmatch(c.data, -1)
	maps := []map[string]string{}

	for _, mgroup := range matches {
		matchMap := map[string]string{}
		for i, match := range mgroup {
			matchMap[names[i]] = match
		}
		maps = append(maps, matchMap)
	}
	return maps
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
