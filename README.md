# go-lru
Simple generic Go LRU cache implementation

## Installation
```
go get github.com/notEpsilon/go-lru
```

## Usage
using the cache is very simple
```
cache := lru.New[int, string](2) // give capacity

cache.Put(1, "fish")

value, err := cache.Get(1) // returns an error if the key wasn't found
if err != nil {
  panic(err)
}

fmt.Println(value) // fish
```
