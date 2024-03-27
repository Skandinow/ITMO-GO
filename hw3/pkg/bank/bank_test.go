package bank

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAccountImpl_Balance(t *testing.T) {
	acc1 := NewAccount(NewRealTime())
	require.Equal(t, acc1.balance, 0)

	acc1.TopUp(10)
	require.Equal(t, acc1.balance, 10)

	acc1.TopUp(100)
	require.Equal(t, acc1.balance, 110)
}

func TestAccountImpl_Operations(t *testing.T) {
	acc1 := NewAccount(NewRealTime())

	for i := 0; i < 5; i++ {
		acc1.TopUp(10)
	}
	for i := 0; i < 6; i++ {
		acc1.Withdraw(10)
	}

	require.Equal(t, len(acc1.operations), 10)
}

func TestAccountImpl_Withdraw(t *testing.T) {
	acc1 := NewAccount(NewRealTime())

	require.False(t, acc1.Withdraw(10))
	require.True(t, acc1.TopUp(10))
	require.True(t, acc1.Withdraw(10))

}

func TestAccountImpl_TopUp(t *testing.T) {
	acc1 := NewAccount(NewMockTime())

	acc1.balance = -100
	require.True(t, acc1.TopUp(99))
	require.True(t, acc1.TopUp(1))
	require.True(t, acc1.TopUp(99))
}

func TestAccountImpl_Statement(t *testing.T) {
	clock := NewMockTime()
	acc1 := NewAccount(clock)
	acc1.TopUp(10)
	require.Equal(t, acc1.Statement(), "2023-03-18 12:34:07 +10 10")
	//fmt.Println(acc1.Statement())

}
