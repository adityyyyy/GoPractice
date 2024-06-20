package main

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, a string) {
		rev, revErr := Reverse(a)
		if revErr != nil {
			return
		}
		doubleRev, doubleErr := Reverse(rev)
		if doubleErr != nil {
			return
		}
		t.Logf(
			"Number of runes: a=%d, rev=%d, doubleRev=%d",
			utf8.RuneCountInString(a),
			utf8.RuneCountInString(rev),
			utf8.RuneCountInString(doubleRev),
		)
		if a != doubleRev {
			t.Errorf("Before: %q, after: %q", a, doubleRev)
		}
		if utf8.ValidString(a) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
