package integer

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(2, 3)
	expect := 5

	if sum != expect {
		t.Errorf("expect %d, but get sum %d", expect, sum)
	}
}
