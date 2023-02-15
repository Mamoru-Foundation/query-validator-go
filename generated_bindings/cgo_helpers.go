// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Tue, 14 Feb 2023 11:16:29 CET.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package generated_bindings

/*
#cgo CFLAGS: -I${SRCDIR}/../packaged/include
#cgo LDFLAGS: -lmamoru_query_validator_go
#cgo darwin,arm64 LDFLAGS: -L${SRCDIR}/../packaged/lib/darwin-arm64
#cgo darwin,amd64 LDFLAGS: -L${SRCDIR}/../packaged/lib/darwin-amd64
#cgo linux,amd64 LDFLAGS: -Wl,--no-as-needed -ldl -lm -L${SRCDIR}/../packaged/lib/linux-amd64
#include <libmamoru_query_validator_go.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"fmt"
	"runtime"
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// allocFfiValidationResultTMemory allocates memory for type C.FfiValidationResult_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocFfiValidationResultTMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfFfiValidationResultTValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfFfiValidationResultTValue = unsafe.Sizeof([1]C.FfiValidationResult_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *FfiValidationResultT) Ref() *C.FfiValidationResult_t {
	if x == nil {
		return nil
	}
	return x.ref5dc80852
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *FfiValidationResultT) Free() {
	if x != nil && x.allocs5dc80852 != nil {
		x.allocs5dc80852.(*cgoAllocMap).Free()
		x.ref5dc80852 = nil
	}
}

// NewFfiValidationResultTRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewFfiValidationResultTRef(ref unsafe.Pointer) *FfiValidationResultT {
	if ref == nil {
		return nil
	}
	obj := new(FfiValidationResultT)
	obj.ref5dc80852 = (*C.FfiValidationResult_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *FfiValidationResultT) PassRef() (*C.FfiValidationResult_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref5dc80852 != nil {
		return x.ref5dc80852, nil
	}
	mem5dc80852 := allocFfiValidationResultTMemory(1)
	ref5dc80852 := (*C.FfiValidationResult_t)(mem5dc80852)
	allocs5dc80852 := new(cgoAllocMap)
	allocs5dc80852.Add(mem5dc80852)

	var cis_error_allocs *cgoAllocMap
	ref5dc80852.is_error, cis_error_allocs = (C._Bool)(x.IsError), cgoAllocsUnknown
	allocs5dc80852.Borrow(cis_error_allocs)

	var cmessage_allocs *cgoAllocMap
	ref5dc80852.message, cmessage_allocs = *(**C.char)(unsafe.Pointer(&x.Message)), cgoAllocsUnknown
	allocs5dc80852.Borrow(cmessage_allocs)

	x.ref5dc80852 = ref5dc80852
	x.allocs5dc80852 = allocs5dc80852
	return ref5dc80852, allocs5dc80852

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x FfiValidationResultT) PassValue() (C.FfiValidationResult_t, *cgoAllocMap) {
	if x.ref5dc80852 != nil {
		return *x.ref5dc80852, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *FfiValidationResultT) Deref() {
	if x.ref5dc80852 == nil {
		return
	}
	x.IsError = (bool)(x.ref5dc80852.is_error)
	x.Message = (*byte)(unsafe.Pointer(x.ref5dc80852.message))
}

// allocSliceRefUint8TMemory allocates memory for type C.slice_ref_uint8_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSliceRefUint8TMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSliceRefUint8TValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfSliceRefUint8TValue = unsafe.Sizeof([1]C.slice_ref_uint8_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *SliceRefUint8T) Ref() *C.slice_ref_uint8_t {
	if x == nil {
		return nil
	}
	return x.refd8594c66
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *SliceRefUint8T) Free() {
	if x != nil && x.allocsd8594c66 != nil {
		x.allocsd8594c66.(*cgoAllocMap).Free()
		x.refd8594c66 = nil
	}
}

// NewSliceRefUint8TRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewSliceRefUint8TRef(ref unsafe.Pointer) *SliceRefUint8T {
	if ref == nil {
		return nil
	}
	obj := new(SliceRefUint8T)
	obj.refd8594c66 = (*C.slice_ref_uint8_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *SliceRefUint8T) PassRef() (*C.slice_ref_uint8_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refd8594c66 != nil {
		return x.refd8594c66, nil
	}
	memd8594c66 := allocSliceRefUint8TMemory(1)
	refd8594c66 := (*C.slice_ref_uint8_t)(memd8594c66)
	allocsd8594c66 := new(cgoAllocMap)
	allocsd8594c66.Add(memd8594c66)

	var cptr_allocs *cgoAllocMap
	refd8594c66.ptr, cptr_allocs = *(**C.uint8_t)(unsafe.Pointer(&x.Ptr)), cgoAllocsUnknown
	allocsd8594c66.Borrow(cptr_allocs)

	var clen_allocs *cgoAllocMap
	refd8594c66.len, clen_allocs = (C.size_t)(x.Len), cgoAllocsUnknown
	allocsd8594c66.Borrow(clen_allocs)

	x.refd8594c66 = refd8594c66
	x.allocsd8594c66 = allocsd8594c66
	return refd8594c66, allocsd8594c66

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x SliceRefUint8T) PassValue() (C.slice_ref_uint8_t, *cgoAllocMap) {
	if x.refd8594c66 != nil {
		return *x.refd8594c66, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *SliceRefUint8T) Deref() {
	if x.refd8594c66 == nil {
		return
	}
	x.Ptr = (*byte)(unsafe.Pointer(x.refd8594c66.ptr))
	x.Len = (uint64)(x.refd8594c66.len)
}

// safeString ensures that the string is NULL-terminated, a NULL-terminated copy is created otherwise.
func safeString(str string) string {
	if len(str) > 0 && str[len(str)-1] != '\x00' {
		str = str + "\x00"
	} else if len(str) == 0 {
		str = "\x00"
	}
	return str
}

// unpackPCharString copies the data from Go string as *C.char.
func unpackPCharString(str string) (*C.char, *cgoAllocMap) {
	allocs := new(cgoAllocMap)
	defer runtime.SetFinalizer(allocs, func(a *cgoAllocMap) {
		go a.Free()
	})

	str = safeString(str)
	mem0 := unsafe.Pointer(C.CString(str))
	allocs.Add(mem0)
	return (*C.char)(mem0), allocs
}

type stringHeader struct {
	Data unsafe.Pointer
	Len  int
}
