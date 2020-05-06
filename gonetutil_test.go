package gonetutil

import "testing"

func TestFqdn(t *testing.T) {
	result := Fqdn()
	t.Logf("FQDN is: %s", result)
	if result == "" {
		t.Error("fqdn() should return either fqdn or unknown, empty string not expected")
	}
}
