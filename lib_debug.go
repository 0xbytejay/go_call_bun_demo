//go:build debug

package main

/*
#cgo linux,amd64 CFLAGS: -fsanitize=address -fno-pic
#cgo linux,amd64 LDFLAGS: -fsanitize=address -fuse-ld=lld -no-pie -fno-pic -fsanitize=address

//Include
#cgo CFLAGS: -I${SRCDIR}/include

//ZIG_LIBS
#cgo linux,amd64 LDFLAGS: -L/root/dev/bun/build/debug/output/libs/linux/x86_64 -Wl,--whole-archive -lbun-debug -Wl,--no-whole-archive

//CPP_LIBS
#cgo linux,amd64 LDFLAGS: -lzstd -lbrotlidec -lhwy -lssl -lbrotlicommon -lmimalloc-asan-debug
#cgo linux,amd64 LDFLAGS: -larchive -lcrypto -lbrotlienc -lz -lcares -ltcc -lls-hpack
#cgo linux,amd64 LDFLAGS: -lsqlite3 -lhdr_histogram_static -ldeflate -ldecrepit -llolhtml

//WEBKIT_LIBS
#cgo linux,amd64 LDFLAGS: -lJavaScriptCore -lWTF -lbmalloc -licui18n -licuuc -licudata -licutu
#cgo linux,amd64 LDFLAGS: -lstdc++ -latomic

#include <stdlib.h>
#include "lbun.h"
*/
import "C"
