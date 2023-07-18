package main

import (
	banks "distributed_election/TCC/Banks"
	manager "distributed_election/TCC/TransactionManager"
	"log"
)

func main() {
	// manager
	// banka, b, c
	var (
		manager = manager.GetSingletonDistributedTransactionManager()
		bankA   = &banks.BankA{
			Balance: 220,
		}
		bankB = &banks.BankB{
			Balance: 50,
		}
		bankC = &banks.BankC{
			Balance: 60,
		}
	)
	err := manager.Try_Confirm_Cancel(bankA, bankB, bankC, -100, 30, 70)
	if err != nil {
		log.Printf("err: %v", err)
	}

	log.Printf("manager: %v", manager)
	log.Printf("bankA: %v", bankA)
	log.Printf("bankB: %v", bankB)
	log.Printf("bankC: %v", bankC)
}
