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

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{name: "multiple words", text: "one two  three", want: "three two one"},
		{name: "single word", text: "solo", want: "solo"},
		{name: "empty", text: "", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.text); got != tt.want {
				t.Errorf("Reverse(%q) = %q, want %q", tt.text, got, tt.want)
			}
		})
	}
}

func TestTagIsStable(t *testing.T) {
	if Tag("a") != Tag("a") || Tag("a") == Tag("b") {
		t.Fatal("Tag must be deterministic and distinct")
	}
}
