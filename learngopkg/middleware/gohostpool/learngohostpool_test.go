package learngopkg

import (
	"errors"
	"sync"
	"testing"

	hostpool "github.com/bitly/go-hostpool"
)

func TestComplicate(t *testing.T) {
	hp := hostpool.NewEpsilonGreedy([]string{"a", "b"}, 0, &hostpool.LinearEpsilonValueCalculator{})
	hostResponse1 := hp.Get()
	hostResponse2 := hp.Get()
	hostResponse1.Mark(errors.New("1"))
	hostResponse2.Mark(errors.New("2"))
	// 并发测试
	const TestNum = 10
	log := make([]string, TestNum)
	counter := &sync.WaitGroup{}
	counter.Add(TestNum)
	for i := 0; i < TestNum; i++ {
		go func(i int) {
			log[i] = hp.Get().Host()
			counter.Done()
		}(i)
	}
	counter.Wait()
	t.Log(log)
}
