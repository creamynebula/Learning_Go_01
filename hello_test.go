package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Oláaaaaa, enfermeira!!\n"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
