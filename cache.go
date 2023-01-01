package cache

import (
	"errors"
)

var usersCache2 = map[string]int{}

type cache struct {
	name_p     string
	value_p    int
	usersCache map[string]int
}

func New() cache {
	i := cache{name_p: "n", value_p: 0}
	i.usersCache = make(map[string]int)
	return i
}

func (user *cache) Set(name string, id int) {
	user.name_p = name
	user.value_p = id
	user.usersCache[name] = id
	usersCache2[name] = id
}

func (user *cache) Get(name string) (int, error) {
	i, ok := user.usersCache[name]
	if ok {
		return i, nil
	}

	return -1, errors.New("key not exist")
}

func (user *cache) Delete(name string) {
	delete(user.usersCache, name)
}
