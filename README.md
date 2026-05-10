![cli-tools banner](assets/banner.svg)

# vhula CLI tools

Personal command-line tools written in Go.

This repository is a learning and practice space inspired by the book
_Powerful Command-Line Applications in Go_. Each tool is intentionally small,
focused, and documented in its own directory.

## Table of Contents

- [Tools](#tools)
- [Requirements](#requirements)
- [Build](#build)
- [Test](#test)
- [Format](#format)
- [Project Layout](#project-layout)
- [Notes](#notes)
- [License](#license)

## Tools

| Tool | Description | Documentation |
| --- | --- | --- |
| `wc` | Counts words, lines, or bytes from standard input. | [wc/README.md](wc/README.md) |
| `todo` | Stores and lists simple todo items in a JSON file. | [todo/README.md](todo/README.md) |

## Requirements

- Go, matching the version declared in each tool's `go.mod`
- `just`, if you want to use the repository task shortcuts

## Build

Build all tools:

```sh
just build
```

This writes binaries to `build/`.

Build an individual tool directly:

```sh
mkdir -p build

cd wc
go build -o ../build/wc main.go

cd ../todo
go build -o ../build/todo cmd/main.go
```

## Test

Run all tests:

```sh
just test
```

Or run the Go tests directly:

```sh
cd wc
go test -v ./...

cd ../todo
go test -v ./...
```

## Format

Format the Go source:

```sh
just fmt
```

## Project Layout

```text
.
├── Justfile        # Build, test, format, and clean tasks
├── LICENSE         # AGPL-3.0 license text
├── README.md       # Repository entry point
├── assets/         # Shared documentation assets
├── todo/           # JSON-backed todo list CLI
│   ├── README.md
│   ├── cmd/
│   │   └── main.go
│   ├── go.mod
│   ├── todo.go
│   └── todo_test.go
└── wc/             # Word/line/byte counter CLI
    ├── README.md
    ├── go.mod
    ├── main.go
    └── main_test.go
```

## Notes

This is a personal collection, not a polished product package. The structure is
kept simple on purpose so each tool can evolve independently as new CLI ideas
are added.

## License

This project is licensed under the GNU Affero General Public License v3.0. See
[LICENSE](LICENSE) for the full text.
