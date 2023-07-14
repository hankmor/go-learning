package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	// noResult()
	// withResult()
	// withResultErr()
	withResultOk()
	// distinguishStdoutAndStderr()
	// grep()
	// setenvDemo()
}

// 直接调用 Cmd 对象的 Run 函数，返回的只有成功和失败，获取不到任何输出的结果。
func noResult() {
	fmt.Println("noResult")
	cmd := exec.Command("ls", "-l", "/var/log/")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("command.Run() failed with %s\n", err)
	}
}

// 有时候我们执行一个命令就是想要获取输出结果，此时你可以调用 Cmd 的 CombinedOutput 函数。
func withResult() {
	fmt.Println("withResult")
	cmd := exec.Command("ls", "-l", "/var/log/")
	// CombinedOutput 函数，只返回 out，并不区分 stdout 和 stderr。
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("command.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}

// shell能执行但是 command 不能执行
func withResultErr() {
	fmt.Println("withResultErr")
	// 出错：ls: /var/log/*.log: No such file or directory
	// 下边的命令执行实际上是等于：ls -l "/var/log/*.log", 将最后的参数当成字符串执行，所以 shell 也不能正确
	// 执行
	cmd := exec.Command("ls", "-l", "/var/log/*.log")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("command.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}

func withResultOk() {
	fmt.Println("withResultErr")
	// 通过 /bin/bash -c 命令将后边的 ls -l /var/log/*.log 作为一个整体来执行，可以实现匹配通配符
	cmd := exec.Command("/bin/bash", "-c", "ls -l /var/log/*.log")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("command.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}

// 上面的写法，无法实现区分标准输出和标准错误，只要换成下面种写法，就可以实现。
func distinguishStdoutAndStderr() {
	fmt.Println("distinguishStdoutAndStderr")
	cmd := exec.Command("ls", "-l", "/var/log/*.log")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	if err != nil {
		log.Fatalf("command.Run() failed with %s\n", err)
	}
}

func grep() {
	fmt.Println("grep")
	c1 := exec.Command("grep", "Usb", "/var/log/wifi.log")
	c2 := exec.Command("wc", "-l")
	c2.Stdin, _ = c1.StdoutPipe()
	c2.Stdout = os.Stdout
	_ = c2.Start()
	_ = c1.Run()
	_ = c2.Wait()
}

func setenvDemo() {
	fmt.Println("setenvDemo")
	os.Setenv("NAME", "demo")
	cmd := exec.Command("echo", os.ExpandEnv("$NAME"))
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("command.Run() failed with %s\n", err)
	}
	fmt.Printf("%s", out)
}
