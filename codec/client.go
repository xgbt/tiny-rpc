package codec

import "io"

type clientCodec struct {
	r io.Reader
	w io.Writer
	c io.Closer
}
