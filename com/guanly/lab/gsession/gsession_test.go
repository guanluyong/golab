package gsession

import (
	"testing"
)

func TestSession(t *testing.T) {
	SM.Set("Key1", "Just test")

	if v := SM.Get("Key1"); v == nil || v.(string) != "Just test" {
		t.Error("Value is", v, ", Not 'Just test'")
	}

	if v := SM.Get("Key2"); v != nil {
		t.Error("Should not have Key2 Session")
	}
}
