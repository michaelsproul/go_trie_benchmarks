package trie_benchmarks

import (
	armon "github.com/armon/go-radix"
	sauerbraten "github.com/sauerbraten/radix"
	"github.com/tchap/go-patricia/patricia"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

// Globals.
var allWords = getText()
var fullPatriciaTrie = makePatriciaTrie(allWords)

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

// Make a Patricia Trie filled with the given words.
func makePatriciaTrie(words []string) *patricia.Trie {
	trie := patricia.NewTrie(patricia.MaxChildrenPerSparseNode(16))

	for _, word := range words {
		trie.Insert(patricia.Prefix(word), len(word))
	}

	return trie;
}

// Main insert benchmark.
func BenchmarkPatriciaInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		makePatriciaTrie(allWords)
	}
}

// Main get benchmark.
func BenchmarkPatriciaGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		results := make([]int, len(allWords))

		for i, word := range allWords {
			results[i] = fullPatriciaTrie.Get(patricia.Prefix(word)).(int)
		}
	}
}

// Main remove bencmark.
func BenchmarkPatriciaInsertRemove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trie := makePatriciaTrie(allWords)

		for _, word := range allWords {
			trie.Delete(patricia.Prefix(word))
		}
	}
}

func BenchmarkSauerbratenInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r := sauerbraten.New()

		for _, word := range allWords {
			r.Set(word, len(word))
		}
	}
}

func BenchmarkArmonInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r := armon.New()

		for _, word := range allWords {
			r.Insert(word, len(word))
		}
	}
}

