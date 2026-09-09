package words

import "testing"

func TestCount(t *testing.T) {
	if got := Count("one two  three"); got != 3 {
		t.Fatalf("Count = %d, want 3", got)
	}
	if got := Count(""); got != 0 {
		t.Fatalf("Count of empty = %d, want 0", got)
	}
}

func TestTagIsStable(t *testing.T) {
	if Tag("a") != Tag("a") || Tag("a") == Tag("b") {
		t.Fatal("Tag must be deterministic and distinct")
	}
}
