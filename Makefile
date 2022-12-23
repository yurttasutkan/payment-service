.PHONY: build clean test serve

build:
	@echo "Compiling source"
	@mkdir -p build
	go build $(GO_EXTRA_BUILD_ARGS) -ldflags "-s -w -X main.version=$(VERSION)" -o build/payment-service cmd/payment-service/main.go

clean:
	@echo "Cleaning up workspace"
	@rm -rf build
	@rm -rf dist

# generate: statics

# statics:
# 	@echo "Generating static files"
# 	statik -src migrations/ -dest internal/ -p migrations -f

test:
	@echo "Running tests"
	go test -p 1 -v ./...

serve: build
	@echo "Starting Payment Service"
	./build/payment-service
