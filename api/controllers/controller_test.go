package controllers

import (
	"testing"
)

func TestGetSubnetIP(t *testing.T) {
	ip := "192.168.0.1"
	prefixSize := 24
	expected := "192.168.0.0"

	result := getSubnetIP(ip, prefixSize)

	if result != expected {
		t.Errorf("getSubnetIP(%q, %d) = %q, expected %q", ip, prefixSize, result, expected)
	}
}
