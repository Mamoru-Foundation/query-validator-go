package query_validator

/*
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"github.com/Mamoru-Foundation/query-validator-go/generated_bindings"
	"unsafe"
)

type Chain byte

const (
	ChainSui Chain = generated_bindings.FfiChainTypeSui
	ChainEvm Chain = generated_bindings.FfiChainTypeEvm
)

func ValidateSql(chain Chain, query string) error {
	result := generated_bindings.FfiValidateSql(generated_bindings.FfiChainTypeT(chain), query)
	defer generated_bindings.FfiDropValidationResult(result)

	result.Deref()

	return makeError(result)
}

func ValidateAssemblyScript(chain Chain, bytes []byte) error {
	cBytes := sliceToFfi(bytes)
	defer freeFfiSlice(cBytes)

	result := generated_bindings.FfiValidateAssemblyScript(generated_bindings.FfiChainTypeT(chain), cBytes)
	defer generated_bindings.FfiDropValidationResult(result)

	result.Deref()

	return makeError(result)
}

func makeError(result generated_bindings.FfiValidationResultT) error {
	if result.IsError {
		message := C.GoString((*C.char)(unsafe.Pointer(result.Message)))

		return errors.New(message)
	} else {
		return nil
	}
}

func sliceToFfi(bytes []byte) generated_bindings.SliceRefUint8T {
	ptr := C.CBytes(bytes)

	return generated_bindings.SliceRefUint8T{
		Ptr: (*byte)(ptr),
		Len: uint64(len(bytes)),
	}
}

func freeFfiSlice(slice generated_bindings.SliceRefUint8T) {
	C.free(unsafe.Pointer(slice.Ptr))
	slice.Free()
}
