# todo

`todo` is a simple task list CLI backed by a JSON file.

## Build

From the repository root:

```sh
just build
```

From this directory:

```sh
mkdir -p ../build
go build -o ../build/todo cmd/main.go
```

## Usage

The examples below assume you are running commands from the repository root
after building the binary.

Print the current task list:

```sh
./build/todo
```

Add a new task:

```sh
./build/todo buy milk
```

Multiple arguments are joined into a single task, so the example above stores
`buy milk`.

## Storage

Tasks are stored in `/.todo.json`.

## Test

```sh
go test -v ./...
```
