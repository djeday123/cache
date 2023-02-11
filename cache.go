package cache

import (
	"fmt"
	"sync"
	"time"
)

type items struct {
	value int
	exp   time.Time
}

type cache struct {
	items map[string]items
	mu    *sync.Mutex
}

const ttl time.Duration = 5 * time.Second

func New() cache {
	item1 := make(map[string]items)
	i := cache{
		items: item1,
		mu:    new(sync.Mutex),
	}

	return i
}

func (user *cache) Set(name string, id int) error {

	user.mu.Lock()
	defer user.mu.Unlock()
	user.items[name] = items{
		value: id,
		exp:   time.Now().Add(ttl),
	}
	return nil
}

func (user *cache) Get(name string) (int, int64, error) {

	user.mu.Lock()
	defer user.mu.Unlock()
	i, ok := user.items[name]

	if ok {
		ttl_check := i.exp.Unix()

		if ttl_check < time.Now().Unix() {
			delete(user.items, name)
			return -1, 0, fmt.Errorf("key %q expired", name)
		} else {
			return i.value, ttl_check, nil
		}
	}

	return -1, 0, fmt.Errorf("key %q not exist", name)
}

func (user *cache) Delete(name string) {
	user.mu.Lock()
	defer user.mu.Unlock()
	delete(user.items, name)
}
