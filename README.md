# fuzzymatch

A Go library that provides a simple API for fuzzy string matching. The implemented algorithm is based on cosine-similarity.

```
package main

import "github.com/mohamadHarith/fuzzymatch"

func main() {
	data := []string{
		"ALADDIN",
		"DARK PHOENIX",
		"GODZILLA: KING OF THE MONSTERS",
		"AQUAMAN",
		"JODOH SYAITAN",
		"ONWARD",
		"THE INVISIBLE MAN",
		"TROLLS WORLD TOUR",
		"CONTAGION! THE BBC FOUR PANDEMIC",
		"MEN IN BLACK: INTERNATIONAL",
		"BLOODSHOT",
		"DORA AND THE LOST CITY OF GOLD",
		"ANGLE HAS FALLEN",
		"IT CHAPTER TWO",
		"CHILD'S PLAY",
		"YESTERDAY",
		"TOY STORY 4",
		"ANABELLE COMESHOME",
		"SPIDER-MAN:FAR FROM HOME",
	}

	matcher := fuzzymatch.New(data,
		fuzzymatch.WithGramSize(2),    // default is 2 (recommended)
		fuzzymatch.WithThreshold(0.5), // default is zero
		fuzzymatch.WithDebug(true),
	)

	res := matcher.Match("contagon pandamix")

	if len(res) > 0 {
		// logic
	}

}
```
