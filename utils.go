package pbcodec

//go:generate protoc --go_out=. pbrpc.proto

import (
	"encoding/binary"
	"io"
)

func WriteRpc(w io.Writer, data []byte) (int, error) {
	sizebuf := make([]byte, 8)
	binary.BigEndian.PutUint64(sizebuf, uint64(len(data)))

	buf := append(sizebuf, data...)
	return w.Write(buf)
}

func ReadRpc(r io.Reader) ([]byte, error) {
	sizebuf := make([]byte, 8)

	_, err := r.Read(sizebuf)
	if err != nil {
		return nil, err
	}

	size := binary.BigEndian.Uint64(sizebuf)

	data := make([]byte, size)
	_, err = r.Read(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
