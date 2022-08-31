package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "\nOláaaaaa, enfermeira!!\n"

	if got != want {
		t.Errorf("got %q want %q", got, want)
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
