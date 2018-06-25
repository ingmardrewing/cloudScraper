package cloudScraper

import "testing"

func TestGetFirstCapturingGroupValue(t *testing.T) {
	c := NewCloudScraper()
	c.(*cloudScraper).data = "asdf assfpa9sf o<Wurst>823<Wurst>123 123 123o"
	c.SetPattern(`st>([^<]+)<`)

	expected := "823"
	actual := c.GetFirstCapturingGroupValue()

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}
