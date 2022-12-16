# Go package for Mamoru Rule Query Validator

## Usage

See `tests/query_validator_test.go`

## Test

```shell
make test
```

## Bench

```shell
make bench
```

## Regenerate bindings

First, install `c-for-go`:
```shell
go install github.com/xlab/c-for-go@latest
```
Generate the bindings:
```shell
make generate-go-bindings
```
