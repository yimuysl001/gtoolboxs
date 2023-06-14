package maputil

import (
	"bytes"
	"encoding/json"
	"strings"
	"sync"
)

var mu sync.RWMutex

// MapCopy returns a copy of the underlying data of the hash map.
func MapCopy[T any](maps map[string]T) map[string]T {
	mu.RLock()
	defer mu.RUnlock()
	data := make(map[string]T, len(maps))
	for k, v := range maps {
		data[k] = v
	}
	return data
}

// Set sets key-value to the hash map.
func Set[T any](maps map[string]T, key string, val T) {
	mu.RLock()
	defer mu.RUnlock()
	maps[key] = val
}

// Set sets key-value to the hash map.
func Get[T any](maps map[string]T, key string) (val T) {
	mu.RLock()
	defer mu.RUnlock()

	if maps != nil {
		val = maps[key]
	}

	return val
}

// Search searches the map with given `key`.
// Second return parameter `found` is true if key was found, otherwise false.
func Search[T any](maps map[string]T, key string) (value T, found bool) {
	mu.RLock()
	defer mu.RUnlock()

	if maps != nil {
		value, found = maps[key]
	}

	return
}

func GetOrSet[T any](maps map[string]T, key string, value T) T {
	if v, ok := Search(maps, key); !ok {

		if maps == nil {
			maps = make(map[string]T)
		}
		maps[key] = value
		return value
	} else {
		return v
	}
}

// GetOrSetFunc returns the value by key,
// or sets value with returned value of callback function `f` if it does not exist
// and then returns this value.
func GetOrSetFunc[T any](maps map[string]T, key string, f func() T) T {
	if v, ok := Search(maps, key); !ok {
		if maps == nil {
			maps = make(map[string]T)
		}
		v = f()
		maps[key] = v
		return v
	} else {
		return v
	}
}

func GetEqualFold[T any](maps map[string]T, key string) (any, bool) {
	for s, t := range maps {
		if strings.EqualFold(s, key) {
			return t, true
		}
	}
	return nil, false
}

func StructTojson(any interface{}) string {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.Encode(any)
	return bf.String()
}
