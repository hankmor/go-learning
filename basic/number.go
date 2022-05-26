package basic

import (
	"bytes"
	"encoding/binary"
)

func TruncateHigh(i int64) int64 {
	var r int64
	for k := int64(10); k <= i; k *= 10 {
		r = i % k
	}
	return r
}

func IntToBytes(n int64) []byte {
	data := n
	bytebuf := bytes.NewBuffer([]byte{})
	err := binary.Write(bytebuf, binary.BigEndian, data)
	if err != nil {
		panic(err)
		return nil
	}
	return bytebuf.Bytes()
}

func BytesToInt(bys []byte) int64 {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	err := binary.Read(bytebuff, binary.BigEndian, &data)
	if err != nil {
		panic(err)
		return 0
	}
	return data
}
