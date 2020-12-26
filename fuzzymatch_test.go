package fuzzymatch_test

import (
	"log"
	"testing"

	"github.com/mohamadHarith/fuzzymatch"
	"github.com/stretchr/testify/require"
)

func TestNormalizeString(t *testing.T) {
	res := fuzzymatch.NormalizeString(` M@#$alay sia@#$`)
	require.Equal(t, "_malay_sia_", res)
}

func TestNew(t *testing.T) {
	m := fuzzymatch.New([]string{"Mississippi"})
	require.NotNil(t, m)
	log.Println(m.GetGramsDictionary())
	log.Println(m.GetVectorMagnitudes())
}

func TestMatch(t *testing.T) {
	m := fuzzymatch.New([]string{"Mississippi", "Malaysia"}, fuzzymatch.WithDebug(true))
	require.NotNil(t, m)
	m.Match("mysia")
}
