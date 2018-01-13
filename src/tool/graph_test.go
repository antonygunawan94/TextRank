package tool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph(t *testing.T) {
	word := Word{
		0,
		[]int{1},
		[]int{2, 5},
		[]int{65, 74},
		"apple",
		0.000005,
		2,
	}

	words := []Word{}
	words = append(words, word)

	sentence := Sentence{
		1,
		"Old apple tree in sunshine.",
	}

	sentences := []Sentence{}
	sentences = append(sentences, sentence)

	graph := Graph{
		sentences,
		words,
	}

	assert.EqualValues(t, "apple", graph.Words[0].Value)
	assert.EqualValues(t, graph.Sentences[0].ID, graph.Words[0].SentenceIDs[0])
}