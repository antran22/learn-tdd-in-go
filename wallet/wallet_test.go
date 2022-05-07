package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnCoin(t *testing.T) {
	t.Run("test underlying value", func(t *testing.T) {
		amount := AnCoin(10)
		assert.EqualValues(t, 10, amount, "underlying value mismatch")
	})

	t.Run("test string representation", func(t *testing.T) {
		amount := AnCoin(20)
		assert.Equal(t, "20 ANC", amount.String(), "invalid string representation")
	})
}

func TestWallet(t *testing.T) {
	t.Run("test deposit normal amount", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(AnCoin(10))

		assertWalletBalance(t, &wallet, AnCoin(10))
	})

	t.Run("test withdraw normal amount", func(t *testing.T) {
		wallet := Wallet{
			balance: AnCoin(23),
		}

		err := wallet.Withdraw(AnCoin(10))

		assert.NoError(t, err, "normal withdrawal must not throw an error")

		assertWalletBalance(t, &wallet, AnCoin(13))
	})

	t.Run("test over-withdrawal", func(t *testing.T) {
		wallet := Wallet{
			6,
		}

		err := wallet.Withdraw(AnCoin(10))

		assert.ErrorIs(t, err, InsufficientFundError, "over-withdrawal must throw an InsufficientFundError")
	})
}

func assertWalletBalance(t testing.TB, wallet *Wallet, expectedBalance AnCoin) {
	t.Helper()

	actualBalance := wallet.Balance()

	assert.Equalf(t, expectedBalance, actualBalance, "balance mismatch, expected %v, actual %v", expectedBalance, actualBalance)
}
