package cache

import (
	"sync"
)

type Cache struct {
	mu   sync.Mutex
	data *sync.Map
}

func NewCache() *Cache {
	return &Cache{
		data: &sync.Map{},
	}
}

func (c *Cache) Set(url string, wordFreq map[string]uint) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data.Store(url, wordFreq)
}

func (c *Cache) Get(url string) (map[string]uint, bool) {
	value, ok := c.data.Load(url)
	if !ok {
		return nil, false
	}
	wordFreq, ok := value.(map[string]uint)
	return wordFreq, ok
}

func (c *Cache) Delete(url string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data.Delete(url)
}
