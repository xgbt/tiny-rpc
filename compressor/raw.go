package compressor

// RawCompressor implements the Compressor interface
type RawCompressor struct {
}

func (RawCompressor) Zip(data []byte) ([]byte, error) {
	return data, nil
}

func (RawCompressor) Unzip(data []byte) ([]byte, error) {
	return data, nil
}
