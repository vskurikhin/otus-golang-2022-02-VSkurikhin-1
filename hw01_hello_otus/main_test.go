package main

import "testing"

func Test_reverse(t *testing.T) {
	for _, c := range []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, ОТУС", "СУТО ,olleH"},
		{"", ""},
	} {
		got := reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func Test_main(_ *testing.T) {
	main()
}
