package banks

import "golang.org/x/xerrors"

/*
	使用Banks来代表每个服务, 实际对balance操作时需要分布式锁
*/

type BankA struct {
	Balance          int64
	RemittanceInput  int64
	RemittanceOutput int64
}

func (p *BankA) RemittanceIn(amount int64) {
	p.Balance += amount
	// 通知分布式事务管理器
	p.Notify("RemittanceIn")
}

func (p *BankA) RemittanceOut(amount int64) error {
	if p.Balance < amount {
		return xerrors.Errorf("BankA 余额不足")
	}

	p.Balance -= amount
	// 通知分布式事务管理器
	p.Notify("RemittanceOut")
	return nil
}

func (p *BankA) Try(amount int64) error {
	if amount > 0 {
		
		p.Notify("Try successfully")
		return nil
	}

	if -amount > p.Balance {
		p.Notify("Try failed")
		return xerrors.Errorf("BankA{Try} 余额不足")
	}

	
	p.Notify("Try successfully")
	return nil
}

func (p *BankA) Confirm(amount int64) error {
	if amount > 0 {
		p.Balance += amount
		p.RemittanceInput = amount
		p.Notify("Confirm successfully")

		return nil
	}
	if -amount > p.Balance {
		p.Notify("Confirm failed")

		return xerrors.Errorf("BankA{Confirm} 余额不足")
	}
	p.Balance += amount
	p.RemittanceOutput = -amount
	p.Notify("Confirm successfully")

	return nil
}

func (p* BankA) Cancel() error {
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

func (p *BankA) Notify(act string) error {
	// connect 分布式事务管理器

	// 发通知ACK to MQ

	return nil
}
