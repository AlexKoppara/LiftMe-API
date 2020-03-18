package liftme

import "testing"

func TestMain(t *testing.T) {
	want := "Hello World"
	if got := main(); got != want {
		t.Errorf("main() = %q, want %q", got, want)
	}
}
