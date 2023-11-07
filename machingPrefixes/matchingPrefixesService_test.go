package machingPrefixes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatcherPrefixes(t *testing.T) {
	// Create a list of sample prefixes
	prefixes := []string{"apple", "app", "banana", "bat", "ball", "cat", "dog"}

	// Build the MatcherDataStore
	matcherDataStore, err := buildMatcherPrefixes(prefixes)
	assert.NoError(t, err)

	// Test individual prefixes
	testPrefix(t, matcherDataStore, "apple", "apple")
	testPrefix(t, matcherDataStore, "apples", "apple")
	testPrefix(t, matcherDataStore, "batman", "bat")
	testPrefix(t, matcherDataStore, "caterpillar", "cat")
	testPrefix(t, matcherDataStore, "xyz", "")

	// Test non-existent prefix
	prefix, err := matcherDataStore.GetMatcherPrefixesRecords(nil, "nonexistent")
	assert.NoError(t, err)
	assert.Equal(t, prefix, "")
}

func testPrefix(t *testing.T, matcherDataStore *getMatcherPrefixesStrore, word, expectedPrefix string) {
	prefix, err := matcherDataStore.GetMatcherPrefixesRecords(nil, word)
	assert.NoError(t, err)
	assert.Equal(t, prefix, expectedPrefix)
}
