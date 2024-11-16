package box

import "reflect"

func Get[T any](box *Box) T {
	box.mu.RLock()
	defer box.mu.RUnlock()

	t := reflect.TypeFor[T]()
	value := box.items[t.String()]
	return value.(T)
}

func GetOrDefault[T any](box *Box, defaultValue T) T {
	box.mu.RLock()
	defer box.mu.RUnlock()

	t := reflect.TypeFor[T]()
	value := box.items[t.String()]
	v, ok := value.(T)

	if !ok {
		return defaultValue
	}

	return v
}

func GetPath[T any](box *Box, path ...string) T {
	box.mu.RLock()
	defer box.mu.RUnlock()

	var value any = box

	for _, key := range path {
		switch v := value.(type) {
		case Box:
			value = v.items[key]
		case *Box:
			value = v.items[key]
		default:
			break
		}
	}

	return value.(T)
}
