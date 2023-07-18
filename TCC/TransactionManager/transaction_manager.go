package manager

import (
	banks "distributed_election/TCC/Banks"

	"golang.org/x/xerrors"
)

type DistributedTransactionManager struct {
}

func NewDistributedTransactionManager() *DistributedTransactionManager {
	return &DistributedTransactionManager{}
}

var singletonDistributedTransactionManager = initSingletonDistributedTransactionManager()

func GetSingletonDistributedTransactionManager() *DistributedTransactionManager {
	return singletonDistributedTransactionManager
}

func initSingletonDistributedTransactionManager() *DistributedTransactionManager {
	return NewDistributedTransactionManager()
}

func (p *DistributedTransactionManager) Try_Confirm_Cancel(bankA *banks.BankA, bankB *banks.BankB, bankC *banks.BankC, amountA, amountB, amountC int64) error {
	err := p.TransactionTrigger(bankA, bankB, bankC, amountA, amountB, amountC)
	if err != nil {
		err = p.Cancel(bankA, bankB, bankC)
		if err != nil {
			return xerrors.Errorf("cancel failed: %w", err)
		}
		return xerrors.Errorf("transaction trigger failed")
	}

	return nil
}

func (p *DistributedTransactionManager) TransactionTrigger(bankA *banks.BankA, bankB *banks.BankB, bankC *banks.BankC, amountA, amountB, amountC int64) error {
	// try
	err := bankA.Try(amountA)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	err = bankB.Try(amountB)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	err = bankC.Try(amountC)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	// confirm
	err = bankA.Confirm(amountA)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	err = bankB.Confirm(amountB)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	err = bankC.Confirm(amountC)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (p *DistributedTransactionManager) Cancel(bankA *banks.BankA, bankB *banks.BankB, bankC *banks.BankC) error {
	err := bankA.Cancel()
	if err != nil {
		// 重试三次
		retryCount := 0
		for bankA.Cancel() != nil {
			retryCount++
			if retryCount == 3 {
				return xerrors.Errorf("failed to cancel transaction: %w", err)
			}
		}
	}

	err = bankB.Cancel()
	if err != nil {
		// 重试三次
		retryCount := 0
		for bankA.Cancel() != nil {
			retryCount++
			if retryCount == 3 {
				return xerrors.Errorf("failed to cancel transaction: %w", err)
			}
		}
	}

	err = bankC.Cancel()
	if err != nil {
		// 重试三次
		retryCount := 0
		for bankA.Cancel() != nil {
			retryCount++
			if retryCount == 3 {
				return xerrors.Errorf("failed to cancel transaction: %w", err)
			}
		}
	}

	return nil
}

func (p *DistributedTransactionManager) AcceptAction(act string, bankType interface{}) {
	switch bankType.(type) {
	case *banks.BankA:
	case *banks.BankB:
	case *banks.BankC:
	}
}
