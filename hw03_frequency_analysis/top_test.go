package hw03frequencyanalysis

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed.
var taskWithAsteriskIsCompleted = true

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive Top10", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"а",         // 8
				"он",        // 8
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"в",         // 4
				"его",       // 4
				"если",      // 4
				"кристофер", // 4
				"не",        // 4
			}
			require.Equal(t, expected, Top10(text))
		} else {
			expected := []string{
				"он",        // 8
				"а",         // 6
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"-",         // 4
				"Кристофер", // 4
				"если",      // 4
				"не",        // 4
				"то",        // 4
			}
			require.Equal(t, expected, Top10(text))
		}
	})
}

func TestSplit(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive split", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			const testSplit2 = `Нога нога нога. нога! cat and dog, one dog,two
                                cats and one man dog`
			expected := []string{
				"нога",
				"нога",
				"нога",
				"нога",
				"cat",
				"and",
				"dog",
				"one",
				"cats",
				"and",
				"one",
				"man",
				"dog",
			}
			require.Equal(t, expected, split(testSplit2))
		} else {
			const testSplit1 = `cat and dog, one dog,two
                                cats and one man dog`
			expected := []string{
				"cat",
				"and",
				"dog,",
				"one",
				"dog,two",
				"cats",
				"and",
				"one",
				"man",
				"dog",
			}
			require.Equal(t, expected, split(testSplit1))
		}
	})
}

func TestMakeBarGraph(t *testing.T) {
	t.Run("no words in empty makeBarGraph", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive makeBarGraph", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			testMakeBarGraph2 := []string{
				"нога",
				"нога",
				"нога",
				"нога",
				"cat",
				"and",
				"dog",
				"one",
				"cats",
				"and",
				"one",
				"man",
				"dog",
			}
			expected := map[string]int{
				"and":  2,
				"cat":  1,
				"cats": 1,
				"dog":  2,
				"man":  1,
				"one":  2,
				"нога": 4,
			}
			require.Equal(t, expected, makeBarGraph(testMakeBarGraph2))
		} else {
			testMakeBarGraph1 := []string{
				"cat",
				"and",
				"dog,",
				"one",
				"dog,two",
				"cats",
				"and",
				"one",
				"man",
				"dog",
			}
			expected := map[string]int{
				"and":     2,
				"one":     2,
				"cat":     1,
				"cats":    1,
				"dog":     1,
				"dog,":    1,
				"dog,two": 1,
				"man":     1,
			}
			require.Equal(t, expected, makeBarGraph(testMakeBarGraph1))
		}
	})
}

func TestMapToEntries(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive mapToEntries", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			testMapToEntries := map[string]int{
				"and":  2,
				"cat":  1,
				"cats": 1,
				"dog":  2,
				"man":  1,
				"one":  2,
				"нога": 4,
			}
			expected := Entries{
				{Word: "нога", Frequency: 4},
				{Word: "and", Frequency: 2},
				{Word: "dog", Frequency: 2},
				{Word: "one", Frequency: 2},
				{Word: "cat", Frequency: 1},
				{Word: "cats", Frequency: 1},
				{Word: "man", Frequency: 1},
			}
			list := mapToEntries(testMapToEntries)
			sort.Sort(list)
			require.Equal(t, expected, list)
		} else {
			testMapToEntries := map[string]int{
				"and":     2,
				"one":     2,
				"cat":     1,
				"cats":    1,
				"dog":     1,
				"dog,":    1,
				"dog,two": 1,
				"man":     1,
			}
			expected := Entries{
				{Word: "and", Frequency: 2},
				{Word: "one", Frequency: 2},
				{Word: "cat", Frequency: 1},
				{Word: "cats", Frequency: 1},
				{Word: "dog", Frequency: 1},
				{Word: "dog,", Frequency: 1},
				{Word: "dog,two", Frequency: 1},
				{Word: "man", Frequency: 1},
			}
			list := mapToEntries(testMapToEntries)
			sort.Sort(list)
			require.Equal(t, expected, list)
		}
	})
}

func TestRankByFrequency(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive rankByFrequency", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			testRankByFrequency := Entries{
				{Word: "нога", Frequency: 4},
				{Word: "and", Frequency: 2},
				{Word: "dog", Frequency: 2},
				{Word: "one", Frequency: 2},
				{Word: "cat", Frequency: 1},
				{Word: "cats", Frequency: 1},
				{Word: "man", Frequency: 1},
			}
			expected := Entries{
				{Word: "нога", Frequency: 4},
				{Word: "and", Frequency: 2},
				{Word: "dog", Frequency: 2},
				{Word: "one", Frequency: 2},
				{Word: "cat", Frequency: 1},
				{Word: "cats", Frequency: 1},
				{Word: "man", Frequency: 1},
			}
			require.Equal(t, expected, rankByFrequency(testRankByFrequency))
		} else {
			testRankByFrequency := Entries{
				{Word: "and", Frequency: 2},
				{Word: "one", Frequency: 2},
				{Word: "cat", Frequency: 1},
				{Word: "cats", Frequency: 1},
				{Word: "dog", Frequency: 1},
				{Word: "dog,", Frequency: 1},
				{Word: "dog,two", Frequency: 1},
				{Word: "man", Frequency: 1},
			}
			expected := Entries{
				{Word: "and", Frequency: 2},
				{Word: "one", Frequency: 2},
				{Word: "cat", Frequency: 1},
				{Word: "cats", Frequency: 1},
				{Word: "dog", Frequency: 1},
				{Word: "dog,", Frequency: 1},
				{Word: "dog,two", Frequency: 1},
				{Word: "man", Frequency: 1},
			}
			require.Equal(t, expected, rankByFrequency(testRankByFrequency))
		}
	})
}
