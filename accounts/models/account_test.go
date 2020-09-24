package models

import (
	"sync"
	"testing"

	"github.com/mcabezas/agileEngine/internal/money"
)

func TestCanReadMultipleBalancesWithoutLocking(t *testing.T) {
	acc := NewAccount("", money.Money{})
	wg := &sync.WaitGroup{}
	wg.Add(10000)
	for i:=0; i<10000; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			_ = acc.Balance()
		}(wg)
	}
	wg.Wait()
}

