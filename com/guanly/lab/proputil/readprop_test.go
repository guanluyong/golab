package proputil

import (
	"testing"
)

func TestGetProps(t *testing.T) {
	m := GetProps("test.properties")
	if m["name"] != "steven" {
		t.Error("field name's value is not steven, is:", m["name"])
	}
}
