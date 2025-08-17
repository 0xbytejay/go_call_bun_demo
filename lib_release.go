//go:build !debug && !rel_debug

package main

/*
#cgo linux,amd64 CFLAGS: -fsanitize=address
#cgo linux,amd64 LDFLAGS: -Wl,--strip-all -Wl,--strip-debug -Wl,--discard-all -fuse-ld=lld -no-pie -fsanitize=address -rdynamic

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
