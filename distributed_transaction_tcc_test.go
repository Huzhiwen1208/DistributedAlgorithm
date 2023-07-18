package main_test

import (
	banks "distributed_election/TCC/Banks"
	manager "distributed_election/TCC/TransactionManager"
	"log"
	"testing"

	"github.com/bytedance/mockey"
	"github.com/stretchr/testify/assert"
	"golang.org/x/xerrors"
)

func TestTry_Confirm_Cancel_should_return_nil_when_has_no_error(t *testing.T) {
	// arrangement
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

	// act
	err := manager.Try_Confirm_Cancel(bankA, bankB, bankC, -100, 30, 70)
	if err != nil {
		log.Printf("err: %v", err)
	}

	// assert
	log.Printf("manager: %v", manager)
	log.Printf("bankA: %v", bankA)
	log.Printf("bankB: %v", bankB)
	log.Printf("bankC: %v", bankC)
}

func TestTry_Confirm_Cancel_should_cancel_when_has_error_in_trying(t *testing.T) {
	// arrangment
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
	// mock
	m := mockey.Mock((*banks.BankA).Try).Return(xerrors.Errorf("mock error")).Build()
	defer m.UnPatch()

	// act
	err := manager.Try_Confirm_Cancel(bankA, bankB, bankC, -100, 30, 70)
	if err != nil {
		log.Printf("err: %v", err)
	}

	// assert
	log.Printf("manager: %v", manager)
	log.Printf("bankA: %v", bankA)
	log.Printf("bankB: %v", bankB)
	log.Printf("bankC: %v", bankC)
	assert.EqualValues(t, 1, m.Times())
}

func TestTry_Confirm_Cancel_should_cancel_when_has_error_in_confirming(t *testing.T) {
	// arrangment
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
	// mock
	m := mockey.Mock((*banks.BankA).Confirm).Return(xerrors.Errorf("mock error")).Build()
	defer m.UnPatch()

	// act
	err := manager.Try_Confirm_Cancel(bankA, bankB, bankC, -100, 30, 70)
	if err != nil {
		log.Printf("err: %v", err)
	}

	// assert
	log.Printf("manager: %v", manager)
	log.Printf("bankA: %v", bankA)
	log.Printf("bankB: %v", bankB)
	log.Printf("bankC: %v", bankC)
	assert.EqualValues(t, 1, m.Times())
}
