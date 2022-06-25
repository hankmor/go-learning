package main

import (
	"bytes"
	"encoding/binary"
	"math"
)

func TruncateHigh(i int64) int64 {
	var r int64
	for k := int64(10); k <= i; k *= 10 {
		r = i % k
	}
	return r
}

func RemainLow(i int64, l int) int64 {
	var r int64
	top := int64(math.Pow10(l - 1))
	if i <= top {
		return i
	}
	for k := int64(10); k <= i; k *= 10 {
		r = i % k
		if k > top {
			break
		}
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
