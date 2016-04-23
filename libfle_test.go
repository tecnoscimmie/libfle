package libfle

import "testing"

func testTools(s string) string {
	return NewFle(s)
}

func TestNewFle(t *testing.T) {
	data := testTools("this is a test")
	if ("this is a test" == data) {
		t.Fatalf("Failed to properly modify the string")
	} else {
		t.Logf("Successfully fle-ed")
	}
}
