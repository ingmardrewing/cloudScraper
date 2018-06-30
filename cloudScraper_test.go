package cloudScraper

import "testing"

func TestGetFirstCapturingGroupValue(t *testing.T) {
	c := NewCloudScraper("asdf assfpa9sf o<Wurst>823<Wurst>123 123 123o")
	c.SetPattern(`st>([^<]+)<`)

	expected := "823"
	actual := c.GetFirstCapturingGroupValue()

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestGetFirstNamedGroupValue(t *testing.T) {
	c := NewCloudScraper("asdf assfpa9sf o<Wurst>823<Wurst>123 123 <Wurst>824<Wurst>123o")
	c.SetPattern(`st>(?P<wurst>[^<]+)<`)

	expected := "823"
	actual := c.GetFirstValueOfGroupNamed("wurst")

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestGetAllNamedGroupValues(t *testing.T) {
	c := NewCloudScraper("asdf assfpa9sf o<Wurst>823</Wurst>123 123 <Wurst>824</Wurst>123o")
	c.SetPattern(`st>(?P<wurst>[^<]+)</`)

	expected := "823"
	actual := c.GetValuesOfGroupsNamed("wurst")[0]

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}

	expected = "824"
	actual = c.GetValuesOfGroupsNamed("wurst")[1]

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}
