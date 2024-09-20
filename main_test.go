package main

import "testing"

func TestHello(t *testing.T) {

	want := "Hello Golang 1111"

	got := hello()

	if want != got {
		t.Fatalf("want %s, got 222 %s\n", want, got)
	}
}
