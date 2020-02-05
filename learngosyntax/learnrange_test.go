package learngosyntax

import (
	"testing"
	"fmt"
)

func TestRangeIndex(t *testing.T) {
	l := []int{1}
	t.Log("for range list []int{1} means [0: 1]")
	t.Log("single variable 'for i := range list':")
	for i := range l {
		t.Log(fmt.Sprintf("i = %v", i))
	}
	t.Log("double variable 'for k, v := range list'")
	for k, v := range l {
		t.Log(fmt.Sprintf("k = %v|v = %v", k, v))
	}
}