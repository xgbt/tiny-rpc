package compressor

import (
	"bytes"
	"io"

	"github.com/golang/snappy"
)

// SnappyCompressor implements the Compressor interface
type SnappyCompressor struct {
}

func (SnappyCompressor) Zip(data []byte) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	w := snappy.NewBufferedWriter(buf)
	defer func() {
		w.Close()
	}()

	_, err := w.Write(data)
	if err != nil {
		return nil, err
	}

	err = w.Flush()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), err
}

func (SnappyCompressor) Unzip(data []byte) ([]byte, error) {
	r := snappy.NewReader(bytes.NewBuffer(data))

	data, err := io.ReadAll(r)
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		return nil, err
	}

	return data, nil
}
