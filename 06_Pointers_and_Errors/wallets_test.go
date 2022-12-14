package main

import "testing"

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {

		wallet := Wallet{}          // declarando uma wallet vazia
		wallet.Deposit(Bitcoin(10)) // balance: Bitcoin(10), em tese

		assertBalance(t, wallet, Bitcoin(10)) // chama t.Errorf se balance != Bitcoin(10)
	})

	t.Run("withdraw with funds", func(t *testing.T) {

		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10)) // balance: Bitcoin(10)

		assertNoError(t, err)                 // se err != nil, chama t.Fatal
		assertBalance(t, wallet, Bitcoin(10)) // chama t.Errorf se balance != Bitcoin(10)
	})

	t.Run("withdraw without sufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100)) // tentando tirar mais que o balance, err != nil

		assertError(t, err, ErrInsufficientFunds) // o teste dá bom se err == ErrInsufficientFunds
		assertBalance(t, wallet, Bitcoin(20))     // e se o saldo tá correto no fim, igual o inicial
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error, but didn't expect one!")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error, but wanted one") //fatal para os testes, não queremos continuar o teste se não tem erro, é um teste do erro
	}

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
