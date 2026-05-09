build: clean test build-wc

build-wc:
    mkdir -p build
    cd wc && go build -o ../build/wc main.go

test:
    cd wc && go test -v ./...

clean:
    rm -rf build

fmt:
    cd wc && go fmt ./...
