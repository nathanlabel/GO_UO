package NetworkServer

import "testing"

func TestHello(t *testing.T) {
	if Hello() == "" {
		t.Error("Hello() should not be nothing")
	}
}