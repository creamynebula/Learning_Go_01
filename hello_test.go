package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chihaya")
	want := "Hello, Chihaya"

	if got != want {
		t.Errorf("Recebemos %q, mas queríamos %q. Baka!", got, want)
	}
}

func TestChihaya(t *testing.T) {
	got := Chihaya()
	//got := "\nOe Kanade!\n"
	want := "\nAyase Chihaya!\n"

	if got != want {
		t.Errorf("Recebemos %q, mas queríamos %q. Baka!", got, want)
	}
}
