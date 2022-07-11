package main

import "testing"

func TestSetAndGet(t *testing.T) {
	if err := CacheSet("xue", []byte("you are stupid")); err != nil {
		t.Fail()
		return
	}
	if v, err := CacheGet("xue"); err != nil {
		t.Fail()
		return
	} else {
		t.Log(string(v))
	}
}
