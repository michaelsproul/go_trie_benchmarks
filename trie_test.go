package trie_benchmarks

import (
	armon "github.com/armon/go-radix"
	sauerbraten "github.com/sauerbraten/radix"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

// Get a list of strings from the file pointed to by $TEST_FILE.
func getText() []string {
	filename := os.Getenv("TEST_FILE")

	if filename == "" {
		panic("$TEST_FILE undefined.")
	}

	contents, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	words := strings.Fields(string(contents))

	return words
}

func BenchmarkSauerbratenInsert(b *testing.B) {
	words := getText()

	for i := 0; i < b.N; i++ {
		r := sauerbraten.New()

		for _, word := range words {
			r.Set(word, len(word))
		}
	}
}

func BenchmarkArmonInsert(b *testing.B) {
	words := getText()

	for i := 0; i < b.N; i++ {
		r := armon.New()

		for _, word := range words {
			r.Insert(word, len(word))
		}
	}
}
