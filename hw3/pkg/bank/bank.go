package bank

import (
	"fmt"
	"time"
)

const (
	TopUpOp OperationType = iota
	WithdrawOp
)

type OperationType int64

type Clock interface {
	Now() time.Time
}

func NewRealTime() *RealClock {
	return &RealClock{}
}

type RealClock struct{}

func (c *RealClock) Now() time.Time {
	return time.Now()
}

type Operation struct {
	OpTime   time.Time
	OpType   OperationType
	OpAmount int
	Balance  int
}

func (o Operation) String() string {
	var format string
	if o.OpType == TopUpOp {
		format = `%s +%d %d`
	} else {
		format = `%s -%d %d`
	}
	return fmt.Sprintf(format, o.OpTime.String()[:19], o.OpAmount, o.Balance)
}

type Account interface {
	TopUp(amount int) bool
	Withdraw(amount int) bool
	Operations() []Operation
	Statement() string
	Balance() int
}

func NewAccount(clock Clock) *AccountImpl {
	return &AccountImpl{balance: 0, operations: []Operation{}, clock: clock}
}

type AccountImpl struct {
	balance    int
	operations []Operation
	clock      Clock
}

func (r *AccountImpl) TopUp(amount int) bool {
	if amount <= 0 {
		return false
	}

	r.makeOperation(r.clock, amount, TopUpOp)

	return true
}

func (r *AccountImpl) Withdraw(amount int) bool {
	if r.balance-amount < 0 || amount <= 0 {
		return false
	}

	r.makeOperation(r.clock, amount, WithdrawOp)

	return true
}

func (r *AccountImpl) makeOperation(clock Clock, amount int, operationType OperationType) {
	if operationType == WithdrawOp {
		r.balance -= amount
	} else {
		r.balance += amount

	}

	r.operations = append(r.operations, Operation{OpTime: clock.Now(),
		OpType: operationType, OpAmount: amount, Balance: r.balance})
}

func (r *AccountImpl) Operations() []Operation {
	return r.operations
}

func (r *AccountImpl) Statement() string {
	s := ""

	for i := 0; i < len(r.operations)-1; i++ {
		s += r.operations[i].String() + "\n"
	}

	if len(r.operations) > 0 {
		s += r.operations[len(r.operations)-1].String()
	}

	return s
}

func (r *AccountImpl) Balance() int {
	return r.balance
}
