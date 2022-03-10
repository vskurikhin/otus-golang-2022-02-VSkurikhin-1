package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

const EMPTY = ""

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if len(s) == 0 {
		return s, nil
	}
	count := uint64(0)
	runes := []rune(s)
	h, t := splitToHeadAndTail(runes, &count)
	result, err := unpackRuneLetter(*h, t, &count)
	return result, err
}

func unpackRuneLetter(r rune, runes []rune, i *uint64) (string, error) {
	if unicode.IsDigit(r) {
		return EMPTY, ErrInvalidString
	}
	result := string(r)
	head, tail := splitToHeadAndTail(runes, i)
	if head != nil {
		if unicode.IsDigit(*head) {
			result = unpackRuneNumber(r, *head)
			head, tail = splitToHeadAndTail(tail, i)
			if nil == head {
				return result, nil
			}
		}
		u, err := unpackRuneLetter(*head, tail, i)
		if err != nil {
			return EMPTY, err
		}
		result += u
	}
	return result, nil
}

func unpackRuneNumber(r rune, p rune) string {
	n, _ := strconv.Atoi(string(p))
	return strings.Repeat(string(r), n)
}

func splitToHeadAndTail(runes []rune, i *uint64) (*rune, []rune) {
	if nil == runes || len(runes) == 0 {
		return nil, nil
	}
	*i++
	if 1 == len(runes) {
		return &runes[0], nil
	}
	return &runes[0], runes[1:]
}
