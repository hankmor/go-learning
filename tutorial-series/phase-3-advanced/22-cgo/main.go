package main

/*
#include <stdio.h>
#include <stdlib.h>

void say_hello(const char* s) {
    printf("Hello form C: %s\n", s);
}
*/
import "C" // 这行代码必须紧跟在注释后面，不能有空行！
import "unsafe"

func main() {
    // Go 字符串不能直接传给 C，需要转换
    s := C.CString("World")
    // C 的内存需要手动释放（Free），Go 的 GC 管不到！
    defer C.free(unsafe.Pointer(s))

    // 调用 C 函数
    C.say_hello(s)
}
