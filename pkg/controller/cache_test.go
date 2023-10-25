package controller

import (
	"testing"
)

func TestCacheSetGetDelete(t *testing.T) {
	Init()
	c := GetCache()

	key := "testKey"
	value := "testValue"

	// Test Set
	c.Set(key, value)

	// Test Get
	got, found := c.Get(key)
	if !found {
		t.Errorf("Expected key %s to be found, but it was not.", key)
	}

	if got != value {
		t.Errorf("Expected value %s, but got %s", value, got)
	}

	// Test Delete
	c.Delete(key)
	_, found = c.Get(key)
	if found {
		t.Errorf("Expected key %s to be deleted, but it was found.", key)
	}
}
