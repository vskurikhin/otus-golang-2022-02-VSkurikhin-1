package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

const EMPTY = ""

const BACKSLASH = rune('\\')

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if len(s) == 0 {
		return s, nil
	}
	runes := []rune(s)
	head, tail := splitToHeadAndTail(runes)
	result, err := unpackRuneLetter(*head, tail)
	return result, err
}

// Закрытая ф-ция которая осуществляет примитивную распаковку строки.
func unpackRuneLetter(current rune, runes []rune) (string, error) {
	if unicode.IsDigit(current) {
		return EMPTY, ErrInvalidString
	}
	result := string(current)
	head, tail := splitToHeadAndTail(runes)
	if head == nil {
		return result, nil
	}
	if BACKSLASH == current {
		var err error
		result, err = unpackRuneBackslash(*head)
		if err != nil {
			return EMPTY, err
		}
		head, tail = splitToHeadAndTailBackslashNumber(&current, head, tail)
		if nil == head {
			return result, nil
		}
	}
	if unicode.IsDigit(*head) {
		result = unpackRuneNumber(current, *head)
		head, tail = splitToHeadAndTail(tail)
		if nil == head {
			return result, nil
		}
	}
	u, err := unpackRuneLetter(*head, tail)
	if err != nil {
		return EMPTY, err
	}
	result += u
	return result, nil
}

// Вспомогательная ф-ция которая проверяет экранируемые символы,
// если это не числовой символ или не братная косая черта то возвращается ошибка.
func unpackRuneBackslash(next rune) (string, error) {
	if !unicode.IsDigit(next) && next != BACKSLASH {
		return EMPTY, ErrInvalidString
	}
	return string(next), nil
}

// Вспомогательная ф-ция которая распаковывает текущий символ на p-раз.
func unpackRuneNumber(current rune, p rune) string {
	n, _ := strconv.Atoi(string(p))
	return strings.Repeat(string(current), n)
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
