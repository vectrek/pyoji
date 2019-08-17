package pyoji

import "testing"

func TestAll(t *testing.T) {
	Reset()
	if Count() != 0 {
		t.Error("start count not zero ")
	}
	Got("foo")
	if Count() != 1 {
		t.Error("didn't get foo")
	}
	Reset()
	if Count() != 0 {
		t.Error("reset: count not zero ")
	}
}