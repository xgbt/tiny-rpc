package serializer

// Serializer is the interface that wraps the basic Marshal and Unmarshal methods.
type Serializer interface {
	Marshal(message interface{}) ([]byte, error)
	Unmarshal(data []byte, message interface{}) error
}
