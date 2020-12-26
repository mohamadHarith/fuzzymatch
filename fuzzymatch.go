package fuzzymatch

import (
	"log"
	"math"
	"regexp"
	"sort"
	"strings"
)

// FuzzyMatcher :
type FuzzyMatcher struct {
	// gramsDictionary : all the strings in the data array are broken down
	// into grams of a fixed size and stored in gramsDictionary map
	// where the key is the gram and the value is another map whose key
	// corresponds to the index of a string in the data array
	// where the gram exists, the value is the count of the gram in that string.
	gramsDictionary map[string]map[int]int
	// gramSize : all the strings in the data array, after normalized will be
	// broken into a fixed size. The default gram size is two.
	gramSize int
	// matchThreshold : the cosine similarity threshold of the matched strings.
	matchThreshold float64
	// originalStrings : stores the data array
	originalStrings []string
	// vectorMagnitudes: stores the vector magnitudes of the normalized strings
	// from the data array. The order of the elements of vectorMagnitudes follows
	// the originalStrings.
	vectorMagnitudes []float64
	// marcherOptions :
	matcherOptions options
}

// NormalizeString : removes special characters and replaces whitespaces,
// start and end of the string with _ .
func NormalizeString(str string) string {
	str = strings.TrimSpace(str)
	str = strings.ToLower(str)
	m, _ := regexp.Compile(`[^a-zA-Z0-9\s]`)
	str = m.ReplaceAllString(str, "")
	str = strings.ReplaceAll(str, " ", "_")
	return "_" + str + "_"
}

// New : takes a parameter of an array of strings from which
// a grams dictionary will be built.
func New(dataArray []string, opts ...Option) *FuzzyMatcher {
	m := new(FuzzyMatcher)
	m.gramsDictionary = make(map[string]map[int]int)
	m.originalStrings = dataArray
	m.vectorMagnitudes = make([]float64, len(dataArray), len(dataArray))
	m.gramSize = 2
	m.matchThreshold = 0.1

	opt := new(options)
	for _, cb := range opts {
		cb(opt)
	}

	m.matcherOptions = *opt

	if m.matcherOptions.gramSize > 0 {
		m.gramSize = m.matcherOptions.gramSize
	}

	if m.matcherOptions.threshold > 0 {
		m.matchThreshold = m.matcherOptions.threshold
	}

	// build grams dictionary
	for i := 0; i < len(dataArray); i++ {
		str := NormalizeString(dataArray[i])

		tempMap := make(map[string]int)

		for j := 0; j < len(str)-m.gramSize+1; j++ {
			gram := str[j : j+m.gramSize]
			if _, ok := m.gramsDictionary[gram]; !ok {
				m.gramsDictionary[gram] = make(map[int]int)
			}
			m.gramsDictionary[gram][i]++
			tempMap[gram]++
		}

		// calculate vector magnitude
		sum := 0
		for _, v := range tempMap {
			sum += (v * v)
		}
		m.vectorMagnitudes[i] = math.Sqrt(float64(sum))
	}

	return m
}

// GetGramsDictionary :
func (f *FuzzyMatcher) GetGramsDictionary() map[string]map[int]int {
	return f.gramsDictionary
}

// GetVectorMagnitudes :
func (f *FuzzyMatcher) GetVectorMagnitudes() []float64 {
	return f.vectorMagnitudes
}

// Match :
type Match struct {
	OrignalString    string
	CosineSimilarity float64
}

// Match :
func (f *FuzzyMatcher) Match(query string) (res []string) {

	query = NormalizeString(query)

	// break query into grams and count them
	gramCount := make(map[string]int)
	for i := 0; i < len(query)-f.gramSize+1; i++ {
		gramCount[query[i:i+f.gramSize]]++
	}

	// vector magnitude of query
	countSum := 0
	var vectorMag float64

	// calculate dot products
	dotProducts := make(map[int]float64) // key is the index of the matched string in the data array

	for k, v := range gramCount {
		countSum += (v * v)
		if i, ok := f.gramsDictionary[k]; ok {
			for idx, count := range i {
				dotProducts[idx] += float64(v) * float64(count)
			}
		}
	}

	vectorMag = math.Sqrt(float64(countSum))

	matches := make([]Match, 0)
	for k, v := range dotProducts {
		cosineSimilarity := v / (vectorMag * f.vectorMagnitudes[k])
		if cosineSimilarity > f.matchThreshold {
			matches = append(matches, Match{
				OrignalString:    f.originalStrings[k],
				CosineSimilarity: cosineSimilarity},
			)
		}
	}

	sort.SliceStable(matches, func(i, j int) bool {
		return matches[i].CosineSimilarity > matches[j].CosineSimilarity
	})

	if f.matcherOptions.debug {
		log.Println("query=> ", query)
		log.Println("matches: ")
	}

	for _, v := range matches {
		if f.matcherOptions.debug {
			log.Printf("match=> %s cosine-similarity=> %f", v.OrignalString, v.CosineSimilarity)
		}
		res = append(res, v.OrignalString)
	}

	return
}
