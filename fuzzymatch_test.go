package fuzzymatch

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNormalizeString(t *testing.T) {
	res := NormalizeString(` M@#$alay sia@#$`)
	require.Equal(t, "_malay_sia_", res)
}

func TestNew(t *testing.T) {
	m := New([]string{"Mississippi"})
	require.NotNil(t, m)
	log.Println(m.gramsDictionary)
	log.Println(m.vectorMagnitudes)
}

func TestMatch(t *testing.T) {
	m := New([]string{"Mississippi", "Malaysia"}, WithDebug(true))
	require.NotNil(t, m)
	m.Match("mysia")
}
