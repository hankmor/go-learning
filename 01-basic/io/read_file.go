package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	// 打开文件
	f, err := os.Open("/etc/passwd")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	// readBytes(f)
	// readBytesBuff(f)

	readLine(f)
}

// 不带缓冲区的读取
func readBytes(f *os.File) {
	// 每次读取的长度
	bs := make([]byte, 1024)
	for {
		// 读取, 返回已经读取的长度，如果为0表示读取完成，也可以通过err == io.EOF来判断读取完成
		i, err := f.Read(bs)
		// if handleerr != nil {
		// 	if handleerr == io.EOF { // 读取完成
		// 		return
		// 	}
		// 	log.Fatal(handleerr)
		// }
		if i == 0 { // 读取完成
			return
		}
		_, err = os.Stdout.Write(bs[:i])
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

// 带缓冲区的读取
func readBytesBuff(f *os.File) {
	// 每次读取的长度
	bs := make([]byte, 1024)
	r := bufio.NewReader(f)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush() // 最后刷新缓冲区
	for {
		i, _ := r.Read(bs)
		if i == 0 { // 读取完成
			return
		}
		_, err := w.Write(bs[:i])
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

func readLine(f *os.File) {
	r := bufio.NewReader(f)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for {
		line, err := r.ReadString('\n') // 按行读取，还有一个更低级的调用 Readline()
		if err != nil {
			return
		}
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Fatal(line)
		}
		w.WriteString(line)
	}
}
