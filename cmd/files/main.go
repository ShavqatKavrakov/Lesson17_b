package main

import (
	"log"

	"github.com/ShavqatKavrakov/Lesson17_b/pkg/types"
	"github.com/ShavqatKavrakov/Lesson17_b/pkg/wallet"
)

func main() {
	srv := &wallet.Service{}
	srv.RegisterAccount("+9920000001")
	srv.Deposit(1, 100_000_00)
	paymet0, _ := srv.Pay(1, "auto", 1_000_00)
	paymet1, _ := srv.Pay(1, "auto", 2_000_00)
	paymet2, _ := srv.Pay(1, "auto", 3_000_00)
	paymet3, _ := srv.Pay(1, "auto", 4_000_00)
	paymet4, _ := srv.Pay(1, "auto", 5_000_00)
	paymet5, _ := srv.Pay(1, "auto", 6_000_00)
	var paymens []types.Payment
	paymens = append(paymens, *paymet0, *paymet1, *paymet2, *paymet3, *paymet4, *paymet5)
	err := srv.HistoryToFile(paymens, "da", 7)
	if err != nil {
		log.Print(err)
	}
}
