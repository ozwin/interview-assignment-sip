package dal

type Store[K comparable, V any] interface {
	Set(key K, value V)
	Get(key K) (V, bool)
}

type KeyValueStore[K comparable, V any] struct {
	store map[K]V
}

func InitializeKeyValueStore[K comparable, V any](capacity int) *KeyValueStore[K, V] {
	return &KeyValueStore[K, V]{
		store: make(map[K]V, capacity),
	}
}

func (kv *KeyValueStore[K, V]) Set(key K, value V) {
	kv.store[key] = value
}

func (kv *KeyValueStore[K, V]) Get(key K) (V, bool) {
	value, exists := kv.store[key]
	return value, exists
}
