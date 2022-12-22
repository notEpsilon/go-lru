package lru_test

import (
	"github.com/notEpsilon/go-lru"
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache, err := lru.New[string, string](2)
	if err != nil {
		t.Errorf("expected `err` to be <nil> but got %q", err.Error())
	}

	cache.Set("name", "ibrahim")
	val, err := cache.Get("name")
	if err != nil {
		t.Errorf("expected `err` to be <nil> but got %q", err.Error())
	}
	if val != "ibrahim" {
		t.Errorf("expected `val` to be %q but got %q", "ibrahim", val)
	}

	cache.Set("age", "21")
	val, err = cache.Get("name")
	if err != nil {
		t.Errorf("expected `err` to be <nil> but got %q", err.Error())
	}
	if val != "ibrahim" {
		t.Errorf("expected `val` to be %q but got %q", "ibrahim", val)
	}
	val, err = cache.Get("age")
	if err != nil {
		t.Errorf("expected `err` to be <nil> but got %q", err.Error())
	}
	if val != "21" {
		t.Errorf("expected `val` to be %q but got %q", "21", val)
	}

	cache.Set("nice", "yes")
	val, err = cache.Get("name")
	if err == nil {
		t.Errorf("expected `err` to be `keyNotFound` error but got <nil>")
	}
	if val != *new(string) {
		t.Errorf("expected `val` to be %q but got %q", *new(string), val)
	}

	val, err = cache.Get("age")
	if err != nil {
		t.Errorf("expected `err` to be <nil> but got %q", err.Error())
	}
	if val != "21" {
		t.Errorf("expected `val` to be %q but got %q", "21", val)
	}

	val, err = cache.Get("nice")
	if err != nil {
		t.Errorf("expected `err` to be <nil> but got %q", err.Error())
	}
	if val != "yes" {
		t.Errorf("expected `val` to be %q but got %q", "yes", val)
	}
}
