# wc

`wc` is a small word-count style utility. It reads from standard input and
prints a single count.

## Build

From the repository root:

```sh
just build
```

From this directory:

```sh
mkdir -p ../build
go build -o ../build/wc main.go
```

## Usage

The examples below assume you are running commands from the repository root
after building the binary.

Count words:

```sh
echo "one two three" | ./build/wc
```

Count lines:

```sh
printf "first\nsecond\n" | ./build/wc -l
```

Count bytes:

```sh
printf "abc" | ./build/wc -b
```

## Flags

| Flag | Description |
| --- | --- |
| `-l` | Count lines instead of words. |
| `-b` | Count bytes instead of words. |

If both `-l` and `-b` are provided, byte counting takes precedence.

## Test

```sh
go test -v ./...
```
