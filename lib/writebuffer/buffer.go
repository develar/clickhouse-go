package writebuffer

import (
	"github.com/valyala/bytebufferpool"
	"io"
)

var bufferPool bytebufferpool.Pool

const InitialSize = 256 * 1024

func New() *WriteBuffer {
	wb := &WriteBuffer{}
	return wb
}

type WriteBuffer struct {
	buffer *bytebufferpool.ByteBuffer
}

func (wb *WriteBuffer) Write(data []byte) (int, error) {
	buffer := wb.buffer
	if buffer == nil {
		buffer = bufferPool.Get()
		wb.buffer = buffer
	}
	return buffer.Write(data)
}

func (wb *WriteBuffer) WriteTo(w io.Writer) (int64, error) {
	buffer := wb.buffer
	if buffer == nil || buffer.Len() == 0 {
		return 0, nil
	}

	n, err := buffer.WriteTo(w)
	wb.Reset()
	return n, err
}

func (wb *WriteBuffer) Bytes() []byte {
	buffer := wb.buffer
	if buffer == nil {
		return []byte{}
	}

	bytes := make([]byte, buffer.Len(), buffer.Len())
	copy(bytes, buffer.Bytes())
	return bytes
}

func (wb *WriteBuffer) len() int {
	buffer := wb.buffer
	if buffer == nil {
		return 0
	}
	return buffer.Len()
}

func (wb *WriteBuffer) Reset() {
	buffer := wb.buffer
	wb.buffer = nil
	bufferPool.Put(buffer)
}
