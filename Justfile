build: clean
    mkdir -p build
    cd wc && go build -o ../build/wc main.go
    cd todo && go build -o ../build/todo cmd/main.go

test:
    cd wc && go test -v ./...
    cd todo && go test -v ./...

clean:
    rm -rf build

fmt:
    cd wc && go fmt ./...
    cd todo && go fmt ./...