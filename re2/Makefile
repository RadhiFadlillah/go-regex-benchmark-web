build:
	@CGO_ENABLED=1 go build -tags re2_cgo -o benchmark-cgo
	@CGO_ENABLED=0 go build -o benchmark-wasm