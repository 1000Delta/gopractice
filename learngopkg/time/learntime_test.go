package learngopkg

import "time"

import "fmt"

import "testing"

func TestTimeAfter(t *testing.T) {
	ch := make(chan int)
	go func(chSend chan<- int) {
		<-time.After(time.Second * 2)
		chSend <- 1
	}(ch)
	t.Log("2s [chan int] test:")
	// 第一次先返回超时
	select {
	case i := <-ch:
		t.Log(fmt.Sprintln(i))
	case <-time.After(time.Second * 1):
		t.Log(fmt.Sprint("timeout"))
	}
	// 第二次先返回数据
	select {
	case i := <-ch:
		t.Log(fmt.Sprintln(i))
	case <-time.After(time.Second * 2):
		t.Log(fmt.Sprint("timeout"))
	}
}
