# Go Calling Bun Static Library Demo

This project is a minimal demonstration of how to call a **Bun** static library from **Go** using **cgo**.
It shows how to:

* Invoke Bun CLI commands from Go.
* Retrieve `JSValue` from `GlobalObject`.
* Convert a `JSValue` to a Go `string`.

---

## Tested Environment

* **OS Kernel**:
  `Linux dev 6.12.38+deb13-amd64 #1 SMP PREEMPT_DYNAMIC Debian 6.12.38-1 (2025-07-16) x86_64 GNU/Linux`
* **Distribution**: Debian 13 (trixie) 
* **Go Version**:
  `go version go1.24.0 linux/amd64`
* **Clang Version**:
  `Debian clang version 19.1.7 (3+b1)`

  > **Note:** Building the Bun static library with Clang `19.1.4` failed due to missing support for certain C++ syntax features.

---

## Environment Setup

Before building, configure the compiler environment variables:

```bash
export CC=clang-19
export CXX=clang++-19
export CGO_LDFLAGS_ALLOW=".*"
```

---

## Build the Demo

```bash
go build -a -v -x -o main main.go
```

This will produce the executable `main`, which links against the Bun static library.

---

## Purpose

This repository only demonstrates:

* Linking a Bun static library into a Go program.
* Executing Bun functionality from Go.
* Performing basic `JSValue` operations.


