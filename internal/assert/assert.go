package assert

import "testing"

func Equal[T comparable](t *testing.T, acutal, expected T) {
	t.Helper()

	if acutal != expected {
		t.Errorf("got: %v want: %v", acutal, expected)
	}
}
