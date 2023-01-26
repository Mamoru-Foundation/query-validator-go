package query_validator

import (
	"C"
	"errors"
	"github.com/Mamoru-Foundation/query-validator-go/generated_bindings"
	"unsafe"
)

type Chain byte

const (
	ChainSui Chain = generated_bindings.FfiChainTypeSui
	ChainEvm Chain = generated_bindings.FfiChainTypeEvm
)

func Validate(chain Chain, query string) error {
	result := generated_bindings.FfiValidate(generated_bindings.FfiChainTypeT(chain), query)
	defer generated_bindings.FfiDropValidationResult(result)

	result.Deref()

	if result.IsError {
		message := C.GoString((*C.char)(unsafe.Pointer(result.Message)))

		return errors.New(message)
	} else {
		return nil
	}
}
