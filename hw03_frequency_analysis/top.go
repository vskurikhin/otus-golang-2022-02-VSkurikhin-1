package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

func Top10(s string) []string {
	m := makeBarGraph(split(s))
	entries := mapToEntries(m)
	return getTop(rankByFrequency(entries), 10)
}

var split = split2

// Для варианта без звёздочки.
// var split1 = func(s string) []string {
//  	return strings.Fields(s)
// }

var patternDropPunctuation = regexp.MustCompile(`^([\wа-яА-Я]+([-,\\.][\wа-яА-Я]+)?)[,.!?]$`)

var patternWord = regexp.MustCompile(`^[\wа-яА-Я]+(-[\wа-яА-Я]+)?$`)

// Вариант co звёздочкой.
var split2 = func(s string) []string {
	result := make([]string, 0)
	for _, word := range strings.Fields(s) {
		matches := patternDropPunctuation.FindSubmatch([]byte(word))
		if len(matches) > 0 {
			result = append(result, strings.ToLower(string(matches[1])))
		} else if patternWord.Match([]byte(word)) {
			result = append(result, strings.ToLower(word))
		}
	}
	return result
}

func makeBarGraph(words []string) map[string]int {
	barGraph := make(map[string]int)
	for _, w := range words {
		barGraph[w]++
	}
	return barGraph
}

type Entry struct {
	Word      string
	Frequency int
}

func (e Entry) Compare(o Entry) int {
	// сравнение Frequency в обратном порядке
	if e.Frequency < o.Frequency {
		return 1
	}
	// по убыванию
	if e.Frequency > o.Frequency {
		return -1
	}
	// а слова в лексикографическом порядке
	if e.Word > o.Word {
		return 1
	}
	if e.Word < o.Word {
		return -1
	}
	return 0
}

type Entries []Entry

func (e Entries) Len() int {
	return len(e)
}

func (e Entries) Less(i, j int) bool {
	return e[i].Compare(e[j]) == -1
}

func (e Entries) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func mapToEntries(barGraph map[string]int) Entries {
	result := make(Entries, 0)
	for word, frequency := range barGraph {
		result = append(result, Entry{Word: word, Frequency: frequency})
	}
	return result
}

func rankByFrequency(list Entries) Entries {
	sort.Sort(list)
	return list
}

func getTop(list Entries, n int) []string {
	result := make([]string, 0, n)
	for i, e := range list {
		result = append(result, e.Word)
		if i > n-2 {
			break
		}
	}
	return result
}
