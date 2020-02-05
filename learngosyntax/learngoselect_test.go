package learngosyntax

import "testing"

// 测试 select 读写顺序
// 随机顺序？	
func TestSelectIOOrder(t *testing.T) {
	chi := make(chan int, 1)
	cho := make(chan int, 1)
	chn := make(chan int, 1)
	cho <- 1
	close(chn)
	select {
	case chi <- 1:
		t.Log("input")
	case <-cho:
		t.Log("output")
	case <-chn:
		t.Log("close")
	default:
		t.Log("default")
	}
}
