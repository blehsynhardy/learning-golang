# Go Learning Project

This repository is a collection of small Go examples and exercises, each in its own folder. It is designed for learning Go fundamentals progressively.

## Prerequisites

- Go 1.18+ installed
- Set `GOPATH` and `GOROOT` as needed (optional with modules)

## Running examples

From project root:

```bash
cd c:\Users\Admin\Documents\golang_project\new_project_go
# run a single example:
cd 8-contact_mngt
go run main.go

# or run all example packages (if no conflicting packages):
go test ./...
# or
# go run ./...
```

## Folder overview

- `6-array_slice/`: array and slice basics
- `7-advance_slice/`: advanced slice operations and behaviors
- `8-contact_mngt/`: simple contact management example with map and slice
- `9-functions/`: function definitions, parameters, results, and closures
- `10-custom-error/`: custom error creation and handling
- `11-defer/`: `defer`, `panic`, and `recover` patterns
- `12-panic_and_recoery/`: explicit panic and recovery examples
- `13-project-math-lib/`: building a math library and package usage
- `14-struc/`: structs and basic data modeling
- `15-method-reciever/`: methods on types and pointer/value receivers
- `16-interfaces/`: interfaces and polymorphism
- `17-stringer-interface/`: Stringer interface implementation (`fmt.Stringer`)
- `18-Generics/`: generic functions and types (Go 1.18+)
- `conditional/`: if/else and switch statements
- `loop/`: for loops and iteration patterns
- `pointers/`: pointer usage and semantics
- `product_mapper/`: example mapping and data transformation

## Notes

- Some folder names contain typos (`recoery`, `reciever`), but code examples are still valid.
- Recommended workflow: open one folder at a time, run `go run main.go`, and read code.

## Contribution

Feel free to add more exercises, fix typos, or improve examples with tests.
