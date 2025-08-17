package main

/*
//Include
#cgo CFLAGS: -I${SRCDIR}/include

#include <stdint.h>
#include <stdlib.h>
#include <string.h>
#include "lbun.h"

typedef bool (*JSCallback)(int16_t);

static bool callJSCallback(JSCallback cb,int16_t i) {
    return cb(i);
}


*/
import "C"
import (
	"context"
	"log"
	"os"
	"runtime"
	"time"
	"unsafe"
)

var jsCallback unsafe.Pointer

var ctx, ready = context.WithCancel(context.Background())

//export BeforeLoadEntryPoint
func BeforeLoadEntryPoint(vm *C.VirtualMachine) {
	log.Println("Bun: BeforeLoadEntryPoint")
	//export func to js here
}

//export RegisterJSCallback
func RegisterJSCallback(globalThis *C.JSGlobalObject, str *C.char, ptr unsafe.Pointer) C.JSValue {
	defer ready()
	log.Printf("Bun: RegisterJSCallback: %s", C.GoString(str))
	jsCallback = ptr
	return 0x7
}

func main() {

	args := os.Args[1:]

	cArgs := make([]*C.char, len(args))
	for i, s := range args {
		cArgs[i] = C.CString(s)
		defer C.free(unsafe.Pointer(cArgs[i]))
	}

	var p runtime.Pinner
	p.Pin(&cArgs[0])
	defer p.Unpin()

	go func() {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
		C.startBunCli((*C.GoSliceHeader)(unsafe.Pointer(&cArgs)))
	}()

	<-ctx.Done()
	go func() {
		var goValue int16 = 0
		for {
			log.Printf("Result from js: %v ", bool(C.callJSCallback((C.JSCallback)(jsCallback), C.int16_t(goValue))))
			goValue += 1
			<-time.After(time.Second * 2)
		}
	}()

	<-make(chan interface{})

}
