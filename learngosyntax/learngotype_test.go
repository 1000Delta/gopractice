package learngosyntax

import "testing"

func TestTypeAboutNil(t *testing.T) {
	// slice
	var slice []int
	t.Logf("[]int before initialize: %v", slice)
	slice = make([]int, 1)
	t.Logf("make([]int, 1) then: %v", slice)

	// struct
	var sxt struct{ Key int }
	t.Logf("struct{Key int} before initialize: %v", sxt)
	sxt = struct{ Key int }{Key: 1}
	t.Logf("struct{Key int}{Key: 1} then: %v", sxt)
}
