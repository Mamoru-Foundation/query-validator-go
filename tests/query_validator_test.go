package tests

import (
	"github.com/Mamoru-Foundation/query-validator-go/query_validator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkSimpleValidation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = query_validator.Validate(query_validator.ChainSui, "SELECT * FROM transactions")
	}
}

func TestValidateValidSuiExpression(t *testing.T) {
	err := query_validator.Validate(query_validator.ChainSui, "SELECT * FROM transactions")

	assert.Nil(t, err)
}

func TestValidateInvalidSuiExpression(t *testing.T) {
	err := query_validator.Validate(query_validator.ChainSui, "SELECT * FROM dummy")

	assert.NotNil(t, err)
}

func TestValidateValidEvmExpression(t *testing.T) {
	err := query_validator.Validate(query_validator.ChainEvm, "SELECT * FROM blocks")

	assert.Nil(t, err)
}

func TestValidateInvalidEvmExpression(t *testing.T) {
	err := query_validator.Validate(query_validator.ChainEvm, "SELECT * FROM dummy")

	assert.NotNil(t, err)
}
