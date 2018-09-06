package simplevm

import "testing"

func TestContext(t *testing.T) {
	ctx := NewContext(123, 3)
	pc := ctx.GetReturnPC()
	if pc != 123 {
		t.Errorf("Program counter should be %v but got %v.", 123, pc)
	}
}
