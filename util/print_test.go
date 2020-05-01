package util

import "testing"

func Test(t *testing.T) {
	var tests = []struct {
		str, want string
	}{
		{"hello", "Util: hello"},
		{"world", "Util: world"},
		{"", "Util: "},
	}
	for _,c := range tests {
		got := Print(c.str)
		if got != c.want {
			t.Errorf("Print(%q) == %q, want %q", c.str, got, c.want)
		}
	}
}
