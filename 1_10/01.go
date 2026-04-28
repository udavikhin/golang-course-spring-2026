package main

import "fmt"

type PaymentProcessor interface {
	Process(amount float64) string
}

type CreditCard struct {
	Currency string
}

func (cc CreditCard) Process(amount float64) string {
	return fmt.Sprintf("Payment (%.2f %s) has been processed", amount, cc.Currency)
}

type CryptoWallet struct {
	Coin string
}

func (cw CryptoWallet) Process(amount float64) string {
	return fmt.Sprintf("Payment (%.2f %s) has been processed", amount, cw.Coin)
}

func main() {
	paymentSystems := []PaymentProcessor{
		CreditCard{Currency: "USD"},
		CryptoWallet{Coin: "BTC"},
	}

	for _, paymentSystem := range paymentSystems {
		fmt.Println(paymentSystem.Process(50))
	}
}
