BINARY_DIR = ./packaged/lib/$$(go env GOOS)-$$(go env GOARCH)
LIB_NAME = libmamoru_query_validator_go.a

dev: build-rust generate-go-bindings test

generate-go-bindings:
	c-for-go cforgo.yaml

build-rust:
	cargo build
	mkdir -p $(BINARY_DIR)
	cp target/debug/$(LIB_NAME) $(BINARY_DIR)/

build-rust-release:
	cargo build --release
	mkdir -p $(BINARY_DIR)
	cp target/release/$(LIB_NAME) $(BINARY_DIR)/

test:
	GODEBUG=cgocheck=2 go test ./tests -v

bench:
	GODEBUG=cgocheck=2 go test -bench=. ./tests

clean:
	cargo clean
	rm -rf generated_bindings/*
