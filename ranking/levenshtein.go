package ranking

import (
	"strings"

	"github.com/renstrom/fuzzysearch/fuzzy"
)

// LevenshteinRanking defines the Levenshtein
// algorithm for string ranking
type LevenshteinRanking struct{}

// Calculate calculates the proximity between the file content and the
// provided sentence using the Levenshtein algorithm
func (lr LevenshteinRanking) Calculate(text string, sentence string) float64 {
	textWords := toWords(text)
	sentenceWords := toWords(sentence)

	var accRanks float64
	for _, sentenceWord := range sentenceWords {
		results := fuzzy.RankFind(sentenceWord, textWords)
		accRanks += maxRank(results)
	}

	return accRanks / float64(len(sentenceWords))
}

// maxRank returns the rank of the best match
// between the sentence and all the text words.
// In case of no rank, 0 is returned, indicating
// the worst rank possible
func maxRank(ranks fuzzy.Ranks) float64 {
	var max = 0.0

	for _, rank := range ranks {
		r := 1.0 - (float64(rank.Distance) / float64(len(rank.Target)))
		if r > max {
			max = r
		}
	}

	return max
}

// toWords split the text in every whitespace and
// line break and return a slice of words
func toWords(text string) []string {
	oneLineText := strings.Replace(text, "\n", " ", -1)
	return strings.Split(oneLineText, " ")
}
