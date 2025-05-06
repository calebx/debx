package main

/*
#cgo CFLAGS: -I/Library/Java/JavaVirtualMachines/zulu-21.jdk/Contents/Home/include -I/Library/Java/JavaVirtualMachines/zulu-21.jdk/Contents/Home/include/darwin
#cgo LDFLAGS: -shared
#include <jni.h>
#include <stdlib.h>

static jstring jx_NewStringUTF(JNIEnv* env, const char* s) {
    return (*env)->NewStringUTF(env, s);
}
*/
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"

	"nothing.com/debx"
)

//export Java_com_example_Hello_SayHello
func Java_com_example_Hello_SayHello(env *C.JNIEnv, clazz C.jclass, jparam C.jstring) C.jstring {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	fmt.Printf(">")
	x := debx.Debx("Hello from Go!")
	fmt.Printf("<")

	cs := C.CString(x)
	defer C.free(unsafe.Pointer(cs))

	// call our wrapper, which invokes (*env)->NewStringUTF(env, cs)
	return C.jx_NewStringUTF(env, cs)
}

func main() {}
