package tests

import (
	"github.com/Mamoru-Foundation/query-validator-go/query_validator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkSimpleValidation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = query_validator.ValidateSql(query_validator.ChainSui, "SELECT * FROM transactions")
	}
}

func TestValidateValidSuiExpression(t *testing.T) {
	err := query_validator.ValidateSql(query_validator.ChainSui, "SELECT * FROM transactions")

	assert.Nil(t, err)
}

func TestValidateInvalidSuiExpression(t *testing.T) {
	err := query_validator.ValidateSql(query_validator.ChainSui, "SELECT * FROM dummy")

	assert.NotNil(t, err)
}

func TestValidateValidEvmExpression(t *testing.T) {
	err := query_validator.ValidateSql(query_validator.ChainEvm, "SELECT * FROM blocks")

	assert.Nil(t, err)
}

func TestValidateInvalidEvmExpression(t *testing.T) {
	err := query_validator.ValidateSql(query_validator.ChainEvm, "SELECT * FROM dummy")

	assert.NotNil(t, err)
}

func TestValidateInvalidAssemblyScript(t *testing.T) {
	err := query_validator.ValidateAssemblyScript(query_validator.ChainEvm, []byte{
		0, 97, 115, 109, 1, 0, 0, 0, 1, 4, 1, 96, 0, 0, 3, 2, 1, 0, 5, 3, 1, 0, 0, 7, 17, 2, 4,
		109, 97, 105, 110, 0, 0, 6, 109, 101, 109, 111, 114, 121, 2, 0, 10, 5, 1, 3, 0, 1, 11, 0,
		36, 16, 115, 111, 117, 114, 99, 101, 77, 97, 112, 112, 105, 110, 103, 85, 82, 76, 18, 46,
		47, 114, 101, 108, 101, 97, 115, 101, 46, 119, 97, 115, 109, 46, 109, 97, 112,
	})

	assert.EqualError(t, err, "DataError(\n"+
		"    WasmExport {\n"+
		"        source: Missing(\n"+
		"            \"__new\",\n"+
		"        ),\n"+
		"        export: \"__new\",\n"+
		"    },\n"+
		")",
	)
}
