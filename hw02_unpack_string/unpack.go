package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

const BACKSLASH = rune('\\')

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if len(s) == 0 {
		return s, nil
	}
	runes := []rune(s)
	head, tail := splitToHeadAndTail(runes)
	var builder strings.Builder
	err := unpackRuneLetter(&builder, *head, tail)
	return builder.String(), err
}

// Закрытая ф-ция которая осуществляет примитивную распаковку строки.
func unpackRuneLetter(b *strings.Builder, current rune, runes []rune) error {
	if unicode.IsDigit(current) {
		return ErrInvalidString
	}
	result := current
	head, tail := splitToHeadAndTail(runes)
	if nil == head {
		b.WriteRune(current)
		return nil
	}
	if BACKSLASH == current {
		if !unicode.IsDigit(*head) && *head != BACKSLASH {
			return ErrInvalidString
		}
		result = *head
		head, tail = splitToHeadAndTailBackslashNumber(&current, head, tail)
		if nil == head {
			b.WriteRune(result)
			return nil
		}
	}
	if unicode.IsDigit(*head) {
		n, err := strconv.Atoi(string(*head))
		if err != nil {
			return err
		}
		WriteRuneNTimes(b, current, n)
		if head, tail = splitToHeadAndTail(tail); nil == head {
			return nil
		}
	} else {
		b.WriteRune(result)
	}
	return unpackRuneLetter(b, *head, tail)
}

// Вспомогательная ф-ция разбивающая массив символов на `голову` и `хвост`,
// при экранировании текущего символа,
// по указателю current записывается экранируемый символ.
func splitToHeadAndTailBackslashNumber(current *rune, head *rune, tail []rune) (*rune, []rune) {
	h, t := splitToHeadAndTail(tail)
	if nil == h {
		return nil, nil
	}
	if unicode.IsDigit(*h) {
		*current = *head
	}
	return h, t
}

// Вспомогательная ф-ция разбивающая массив символов на `голову` и `хвост`.
func splitToHeadAndTail(runes []rune) (*rune, []rune) {
	if nil == runes || len(runes) == 0 {
		return nil, nil
	}
	if len(runes) == 1 {
		return &runes[0], nil
	}
	return &runes[0], runes[1:]
}

func WriteRuneNTimes(b *strings.Builder, r rune, n int) {
	for i := 0; i < n; i++ {
		b.WriteRune(r)
	}
}
