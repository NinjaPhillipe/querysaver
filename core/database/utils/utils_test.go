package utils

import "testing"

func TestConcatArrayInt(t *testing.T) {
	nbs := []int{1, 2, 3, 4}
	expected := "1,2,3,4"
	actual := ConcatArrayInt(nbs, ",")
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
