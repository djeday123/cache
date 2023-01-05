# cache
Golang in-memory cache

#### Introduction

A simple, semantic and developer-friendly golang package for datetime

If you think
it is helpful, please give me a star

[github.com/djeday123/cache](https://github.com/golang-module/carbon "github.com/djeday123/cache")

[gitee.com/djeday123/cache](https://gitee.com/golang-module/carbon "gitee.com/djeday123/cache")

#### Installation

##### Go version >= 1.18 (recommend)

```go
// By github
go get -u github.com/djeday123/cache

import (
	"github.com/djeday123/cache"
)
```

#### Usage and example

```go
##### Create cache instance

cache := cache.New()

##### Example of set key (string) and value (int) in cache
cache.Set("userId", 42)

##### Example of get value of the key in cache
userId, err := cache.Get("userId")
if err != nil {
  fmt.Printf("key userId doesnt exist \n")
} else {
  fmt.Println(userId)
}

##### Example of delete row by the key in cache
cache.Delete("userId")
```
