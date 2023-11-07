package machingPrefixes

import (
	"sync"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type GetMatcherPrefixesService interface {
	GetMatcherPrefixesRecords(c *gin.Context, word string) (string, error)
}

// getCustomerPlatformService represents the getCustomerPlatformService for this application
// swagger:model
type getMatcherPrefixesStrore struct {
	Mu       sync.RWMutex // Mutex for concurrent access
	Children map[rune]*getMatcherPrefixesStrore
	End      bool
	fx.In
}

func NewGetMatchingPrefixesService() GetMatcherPrefixesService {
	return &getMatcherPrefixesStrore{
		Mu:       sync.RWMutex{},
		Children: make(map[rune]*getMatcherPrefixesStrore),
		End:      false,
	}
}

// This function pre-processes the list of prefixes and constructs the MatcherPrefixes in parallel using goroutines
func buildMatcherPrefixes(prefixes []string) (*getMatcherPrefixesStrore, error) {
	var wg sync.WaitGroup
	root := &getMatcherPrefixesStrore{Mu: sync.RWMutex{}, Children: make(map[rune]*getMatcherPrefixesStrore)}

	for _, prefix := range prefixes {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			root.Insert(p)
		}(prefix)
	}

	wg.Wait()
	return root, nil
}

// Insert a word into the MatcherDataStore
func (matcherDataStore *getMatcherPrefixesStrore) Insert(word string) {
	matcherDataStore.Mu.Lock()
	defer matcherDataStore.Mu.Unlock()

	node := matcherDataStore
	for _, ch := range word {
		if node.Children[ch] == nil {
			node.Children[ch] = &getMatcherPrefixesStrore{Mu: sync.RWMutex{}, Children: make(map[rune]*getMatcherPrefixesStrore)}
		}
		node = node.Children[ch]
	}
	node.End = true
}

// GetCustomerDataPlatformRecords...gets the Users records from dynamoDB
func (getCustomerPlatformService *getMatcherPrefixesStrore) GetMatcherPrefixesRecords(c *gin.Context, word string) (string, error) {
	var err error
	getCustomerPlatformService.Mu.RLock()
	defer getCustomerPlatformService.Mu.RUnlock()

	node := getCustomerPlatformService
	var prefix string
	for _, ch := range word {
		if node.Children[ch] != nil {
			prefix += string(ch)
			node = node.Children[ch]
		} else {
			break
		}
	}
	return prefix, err
}
