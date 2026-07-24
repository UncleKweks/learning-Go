# learning-Go

A collection of small Go programs written while learning the language. Each mini folder is a self-contained exercise focused on a particular concept.

## Contents

- [Basics](Basics/main.go) — entry point with basic Go syntax examples (loops, `fmt.Println`).
- [simpleContactManagementSystem](simpleContactManagementSystem/main.go) — an in-memory contact manager demonstrating structs, slices, and maps. Supports adding contacts, listing them, and looking up a contact by name.
- [saveMathLib](saveMathLib/main.go) — custom error types and safe math operations demonstrating Go's error handling patterns.
- [bankAccountManagement](bankAccountManagement/main.go) — a bank account system demonstrating struct embedding and interfaces, with savings and overdraft account variants.
- [payRollProcessor](payRollProcessor/main.go) — a payroll processor demonstrating interfaces and polymorphism across salaried, hourly, and commissioned employees.
- [pingPonger](pingPonger/main.go) — a ping-pong goroutine exercise demonstrating channels, `select`, and `context` cancellation.
- [fileDownloader](fileDownloader/main.go) — a concurrent file downloader demonstrating goroutines, channels, and `sync.WaitGroup` with a bounded worker limiter; downloads a list of URLs in parallel and reports size/duration per file.

## Requirements

- Go 1.26+

## Running

```bash
go run ./Basics
go run ./simpleContactManagementSystem
go run ./saveMathLib
go run ./bankAccountManagement
go run ./payRollProcessor
go run ./pingPonger
go run ./fileDownloader
```
