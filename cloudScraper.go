package cloudScraper

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

func NewCloudScraper() CloudScraper {
	cs := new(cloudScraper)
	cs.timeout = time.Duration(10 * time.Second)
	return cs
}

type CloudScraper interface {
	DownloadData() int
	GetFirstCapturingGroupValue() string
	SetTimeout(time.Duration)
	Data() string
	SetPattern(string)
	SetUrl(string)
}

type cloudScraper struct {
	url     string
	data    string
	timeout time.Duration
	pattern *regexp.Regexp
}

// Sets the the timout for the http request in seconds
func (c *cloudScraper) SetTimeout(timeout time.Duration) {
	c.timeout = timeout
}

// Sets the url of the page to be inspected
func (c *cloudScraper) SetUrl(url string) {
	c.url = url
}

// Sets the regex pattern to use for
// the data extraction. Must contain
// a capturing group for this to work
func (c *cloudScraper) SetPattern(pattern string) {
	if len(pattern) == 0 {
		log.Fatalln("Pattern must not be an empty string")
	}
	c.pattern = regexp.MustCompile(pattern)
}

// compiles the given pattern and returns
// the value of the first capturing group
// defined within the pattern, if any.
// Retuns an empty string if nothing is found
func (c *cloudScraper) GetFirstCapturingGroupValue() string {
	matches := c.pattern.FindStringSubmatch(c.data)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// Loads the data from a target server and return the http
// status code as int
func (c *cloudScraper) DownloadData() int {
	if len(c.url) == 0 {
		log.Fatalln("No url given")
	}
	client := http.Client{
		Timeout: c.timeout,
	}

	resp, err := client.Get(c.url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	c.data = string(responseData)
	return resp.StatusCode
}

// Returns the downloaded data, if any
func (c *cloudScraper) Data() string {
	return c.data
}
