package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

// func NewDecoder(r io.Reader) *Decoder
// func NewEncoder(w io.Writer) *Encoder
func TestStreaming(t *testing.T) {
	s := `{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`
	var b bytes.Buffer
	b.Write([]byte(s))
	// dec := json.NewDecoder(os.Stdin)
	dec := json.NewDecoder(&b)
	enc := json.NewEncoder(os.Stdout)
	for {
		var v map[string]interface{}
		// 解码输入的 json 到 map
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}
		// 打印map
		for k := range v {
			fmt.Printf("key: %v, value: %v\n", k, v[k])
		}
		// 将解码后的 map 写入 os.Stdout 输出流中
		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}
	}
}
