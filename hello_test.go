package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("Dizendo oi pros membros do Karuta Club.", func(t *testing.T) {
		got := Hello("Chihaya")
		want := "Olá, Chihaya!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Olá, enfermeira! Quando chamarem a função com nome vazio.", func(t *testing.T) {
		got := Hello("")
		want := "Olá, enfermeira!"

		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
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
