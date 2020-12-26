package fuzzymatch

type options struct {
	gramSize  int
	threshold float64
	debug     bool
}

// Option :
type Option func(*options)

// WithGramSize :
func WithGramSize(size int) Option {
	return func(o *options) {
		o.gramSize = size
	}
}

// WithThreshold :
func WithThreshold(threshold float64) Option {
	return func(o *options) {
		o.threshold = threshold
	}
}

// WithDebug :
func WithDebug(debug bool) Option {
	return func(o *options) {
		o.debug = debug
	}
}
