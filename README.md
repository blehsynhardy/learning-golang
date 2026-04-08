# Go Learning Project

A collection of Go exercises, examples, and small projects organized by topic. Each folder contains a self-contained example to help learn Go syntax, features, and idioms.

## Requirements

- Go 1.25 or newer
- No special environment setup is required for module-enabled projects

## How to run examples

### Root examples

From the repository root:

```powershell
cd C:\Users\Admin\Documents\golang_project\new_project_go
# Run a single example folder:
cd 8-contact_mngt
go run main.go
```

### Run directly from the root

If the example folder contains a single `main.go`, you can run it directly from the repo root:

```powershell
go run ./8-contact_mngt
```

### Module examples

Some folders use nested Go modules and must be run from their own directory:

```powershell
cd 29-go-modules
go run ./cmd/colorer/
```

```powershell
cd 29-go-modules-ext-libraries
go run ./cmd/colorer/
```

## Example folders

- `6-array_slice/`: arrays, slices, and slice behavior
- `7-advance_slice/`: advanced slice operations and memory behavior
- `8-contact_mngt/`: contact management example with maps and slices
- `9-functions/`: functions, multiple return values, and closures
- `10-custom-error/`: custom errors and error handling
- `11-defer/`: `defer` usage, cleanup, and panic handling
- `12-panic_and_recoery/`: panic recovery examples
- `13-project-math-lib/`: packaging and using a math library
- `14-struc/`: struct definitions and basic data modeling
- `15-method-reciever/`: methods on value and pointer receivers
- `16-interfaces/`: using interfaces for polymorphism
- `17-stringer-interface/`: implementing `fmt.Stringer`
- `18-Generics/`: generic functions and types
- `19-project-payroll/`: payroll example with structs and business logic
- `20-composition/`: composition patterns with structs
- `21-embedding-with-inheritance/`: struct embedding and inheritance-like behavior
- `22-project-bank-acct-mangt/`: bank account management example
- `23-strings/`: string handling and utilities
- `24-strings-formatting/`: formatting strings with `fmt`
- `25-go-unicode/`: Unicode and rune handling
- `26-go-regex/`: regular expressions in Go
- `27-text-template/`: text/template usage
- `28-project-config-parser/`: config file parser example
- `29-go-modules/`: Go modules and package organization
- `29-go-modules-ext-libraries/`: Go modules with external dependencies and nested packages
- `30-intro-to-go-routines/`: goroutine basics and concurrent execution
- `31-waitgroups/`: `sync.WaitGroup` for coordinating goroutines
- `32-go-channels/`: channel communication basics
- `33-buffered-channel/`: buffered channels and capacity behavior
- `34-closing-channels/`: closing channels and range semantics
- `35-project-ping-pong/`: ping-pong concurrency example
- `36-project-concurent-file-downloader/`: concurrent file downloader example
- `37-mutex/`: mutex synchronization and race prevention
- `conditional/`: if/else and switch statements
- `loop/`: for loops and iteration patterns
- `pointers/`: pointer usage and semantics
- `product_mapper/`: example mapping and data transformation

## Notes

- Some folder names include typos. The code itself is the primary learning focus.
- Use `go test ./...` from the module root only when the folder is intended as a module or when package names do not conflict.

## Contribution

Feel free to add more examples, improve the README, fix typos, or add tests for the existing exercises.
