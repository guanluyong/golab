package dbconfig

import (
	"testing"
)

func TestDbKindInvaild(t *testing.T) {
	var k DbKind = 0
	if k.String() != "invalid" {
		t.Error("You got an error.")
	}
	t.Log("Yes, It's right")
}
