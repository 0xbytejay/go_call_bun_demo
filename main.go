package main

/*
#cgo linux,amd64 CFLAGS: -fsanitize=address
#cgo linux,amd64 LDFLAGS: -fsanitize=address -Wl,--strip-all -Wl,--strip-debug -Wl,--discard-all -fuse-ld=lld -no-pie -fsanitize=address

//Include
#cgo CFLAGS: -I${SRCDIR}/include

//ZIG_LIBS
#cgo linux,amd64 LDFLAGS: -L${SRCDIR}/libs/linux/x86_64 -Wl,--whole-archive -lbun -Wl,--no-whole-archive

//CPP_LIBS
#cgo linux,amd64 LDFLAGS: -lzstd -lbrotlidec -lhwy -lssl -lbrotlicommon
#cgo linux,amd64 LDFLAGS: -larchive -lcrypto -lbrotlienc -lz -lcares -ltcc -lls-hpack
#cgo linux,amd64 LDFLAGS: -lsqlite3 -lhdr_histogram_static -ldeflate -ldecrepit -llolhtml

//WEBKIT_LIBS
#cgo linux,amd64 LDFLAGS: -lJavaScriptCore -lWTF -lbmalloc -licui18n -licuuc -licudata -licutu
#cgo linux,amd64 LDFLAGS: -lstdc++ -latomic

#include <stdlib.h>
#include "lbun.h"
*/
import "C"
import (
	"log"
	"runtime"
	"time"
	"unsafe"
)

func main() {

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	args := []string{"run", "/root/dev/go_bun_demo/index.js"}

	cArgs := make([]*C.char, len(args))
	for i, s := range args {
		cArgs[i] = C.CString(s)
		defer C.free(unsafe.Pointer(cArgs[i]))
	}

	var p runtime.Pinner
	p.Pin(&cArgs[0])
	defer p.Unpin()

	go C.startBunCli((*C.GoSlice)(unsafe.Pointer(&cArgs)))

	<-time.After(time.Millisecond * 5000)

	var globalThis = C.VirtualMachine_getMainThreadVMGlobalObject()

	var globalThisValue = C.JSGlobalObject_toJSValue(globalThis)

	var testCustomValue = C.JSValue_get(globalThis, globalThisValue.value, C.CString("testCustomValue"))

	var testCustomValueStrPtr = C.JSValue_toString(globalThis, testCustomValue.value)

	defer C.free(unsafe.Pointer(testCustomValueStrPtr))

	log.Printf("String Value From JS: %s", C.GoString(testCustomValueStrPtr))

	// FIXME: Random Crash
	// var testFuncValue = C.JSValue_get(globalThis, globalThisValue.value, C.CString("testFunc"))
	// if testFuncValue.is_err {
	// 	return
	// }

	// go C.JSValue_call(globalThis, globalThisValue.value, testFuncValue.value, nil)

	<-make(chan interface{})

}
