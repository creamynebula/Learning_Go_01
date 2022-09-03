package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("Dizendo oi pros membros do Karuta Club.", func(t *testing.T) {
		got := Hello("Chihaya", "")
		want := "Olá, Chihaya!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Olá, enfermeira! Quando chamarem a função com nome vazio.", func(t *testing.T) {
		got := Hello("", "")
		want := "Olá, enfermeira!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Amèlie", "French")
		want := "Bonjour, Amèlie!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Hungarian, Bulgarian or Romenian", func(t *testing.T) {
		got := Hello("Adyelya", "Bulgarian")
		want := "Sua linguagem é doidona, Adyelya!"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Recebemos %q, mas queríamos %q. Baka!", got, want)
	}
}
