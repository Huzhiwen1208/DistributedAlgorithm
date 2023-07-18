package banks

import "golang.org/x/xerrors"

type BankB struct {
	Balance          int64
	RemittanceInput  int64
	RemittanceOutput int64
}

func (p *BankB) RemittanceIn(amount int64) {
	p.Balance += amount
	// 通知分布式事务管理器
	p.Notify("RemittanceIn")
}

func (p *BankB) RemittanceOut(amount int64) error {
	if p.Balance < amount {
		return xerrors.Errorf("BankB 余额不足")
	}

	p.Balance -= amount
	// 通知分布式事务管理器
	p.Notify("RemittanceOut")
	return nil
}

func (p *BankB) Notify(act string) error {
	// connect 分布式事务管理器

	// 发通知ACK to MQ

	return nil
}

func (p *BankB) Try(amount int64) error {
	if amount > 0 {
		
		p.Notify("Try successfully")
		return nil
	}

	if -amount > p.Balance {
		p.Notify("Try failed")
		return xerrors.Errorf("BankB{Try} 余额不足")
	}

	p.Notify("Try successfully")
	return nil
}

func (p *BankB) Confirm(amount int64) error {
	if amount > 0 {
		p.Balance += amount
		p.RemittanceInput = amount
		p.Notify("Confirm successfully")

		return nil
	}
	if -amount > p.Balance {
		p.Notify("Confirm failed")

		return xerrors.Errorf("BankB{Confirm} 余额不足")
	}
	p.Balance += amount
	p.RemittanceOutput = -amount
	p.Notify("Confirm successfully")

	return nil
}

func (p *BankB) Cancel() error {
	if p.RemittanceInput != 0 && p.Balance >= p.RemittanceInput {
		p.Balance -= p.RemittanceInput
		p.Notify("Cancel successfully")

		return nil
	}

	if p.RemittanceOutput != 0 {
		p.Balance += p.RemittanceOutput
		p.Notify("Cancel successfully")

		return nil
	}

	return nil
}
