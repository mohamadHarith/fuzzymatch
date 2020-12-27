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
		fuzzymatch.WithDebug(true),    // default is false
	)

	res := matcher.Match("an come home")

	if len(res) > 0 {
		// logic
	}

}
```

## Sample Results

```
2020/12/27 02:58:59 query=>  an come home
2020/12/27 02:58:59 ANABELLE COMESHOME 0.780013
2020/12/27 02:58:59 SPIDER-MAN:FAR FROM HOME 0.504715
```

```
2020/12/27 12:37:06 query=>  contagon pandamix
2020/12/27 12:37:06 CONTAGION! THE BBC FOUR PANDEMIC 0.559017
```

```
2020/12/27 12:39:10 query=>  the invisible man
2020/12/27 12:39:10 THE INVISIBLE MAN 1.000000
```

```
2020/12/27 12:44:14 query=>  zbyderman
2020/12/27 12:44:14 AQUAMAN 0.335410
2020/12/27 12:44:14 SPIDER-MAN:FAR FROM HOME 0.316228
2020/12/27 12:44:14 THE INVISIBLE MAN 0.212132
```

## TODOS

- [ ] Add more testcases.
- [ ] Add benchmarks.
- [ ] Improve debug log output.
- [ ] Clean up codes.
- [ ] Optimize algorithm.

## References and Inspirations

1. https://engineering.continuity.net/cosine-similarity/
2. https://glench.github.io/fuzzyset.js/
