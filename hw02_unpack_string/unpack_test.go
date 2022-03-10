package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		// uncomment if task with asterisk completed
		// {input: `qwe\4\5`, expected: `qwe45`},
		// {input: `qwe\45`, expected: `qwe44444`},
		// {input: `qwe\\5`, expected: `qwe\\\\\`},
		// {input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func TestUnpackRuneLetterInvalidString(t *testing.T) {
	invalidStrings := []string{"1a", "3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			count := uint64(0)
			r := ([]rune(tc))[0]
			runes := ([]rune(tc))[1 : len(tc)-1]
			_, err := unpackRuneLetter(r, runes, &count)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func TestUnpackRuneLetterOneCorrectRune(t *testing.T) {
	strings := []string{"a", "б"}
	for _, tc := range strings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			count := uint64(1)
			r := ([]rune(tc))[0]
			runes := []rune(nil)
			_, err := unpackRuneLetter(r, runes, &count)
			require.Falsef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
			require.Equal(t, uint64(1), count)
		})
	}
}

func TestUnpackRuneLetter(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		count    uint64
	}{
		{input: "a2", expected: "aa", count: 1},
		{input: "b3c", expected: "bbbc", count: 2},
		{input: "б4вгд", expected: "ббббвгд", count: 4},
		{input: "\n5", expected: "\n\n\n\n\n", count: 1},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			count := uint64(0)
			r := ([]rune(tc.input))[0]
			var runes []rune
			if len(tc.input) > 1 {
				runes = ([]rune(tc.input))[1:]
			} else {
				runes = nil
			}
			result, err := unpackRuneLetter(r, runes, &count)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
			require.Equal(t, tc.count, count)
		})
	}
}

func TestSplitToHeadAndTail(t *testing.T) {
	a := 'a'
	ab := []rune("b")
	tests := []struct {
		input     []rune
		expected1 *rune
		expected2 []rune
		expected3 uint64
		count     uint64
	}{
		{input: nil, expected1: nil, expected2: nil, expected3: 0, count: uint64(0)},
		{input: []rune(""), expected1: nil, expected2: nil, expected3: 0, count: uint64(0)},
		{input: []rune("a"), expected1: &a, expected2: nil, expected3: 1, count: uint64(0)},
		{input: []rune("ab"), expected1: &a, expected2: ab, expected3: 1, count: uint64(0)},
	}

	for _, tc := range tests {
		r, runes := splitToHeadAndTail(tc.input, &tc.count)
		if nil == tc.expected1 {
			require.Nil(t, r)
		} else {
			require.NotNil(t, r)
			require.EqualValues(t, *r, *tc.expected1)
		}
		if nil == tc.expected2 {
			require.Nil(t, runes)
		} else {
			require.NotNil(t, runes)
			require.EqualValues(t, runes, tc.expected2)
		}
		require.EqualValues(t, tc.count, tc.expected3)
	}
}
