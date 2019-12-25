package connector

type inMemoryKeyValue struct {
	data map[string]string
}

type InMemoryKeyValue interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

func NewInMemoryKeyValue() InMemoryKeyValue {
	return &inMemoryKeyValue{data: make(map[string]string)}
}

func (i *inMemoryKeyValue) Get(key string) (string, error) {
	return i.data[key], nil
}

func (i *inMemoryKeyValue) Set(key string, value string) error {
	i.data[key] = value
	return nil
}
