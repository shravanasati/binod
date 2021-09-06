package stringutil

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"test-reverse1", "haha", "ahah"},
		{"test-reverse2", "olleh", "hello"},
		{"test-reverse3", "reverse", "esrever"},
		{"test-reverse4", "marvel", "levram"},
		{"test-reverse5", "this is a pretty long sentence", "ecnetnes gnol ytterp a si siht"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseString(tt.arg); got != tt.want {
				t.Errorf("ReverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
