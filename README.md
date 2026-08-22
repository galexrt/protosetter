# protosetter
Welcome to the protosetter project!

## Overview
protosetter is a linter developed specifically for Go programmers working with nested `protobuf` types.\
It's designed to aid developers in preventing `invalid memory address or nil pointer dereference` errors arising from direct access of nested `protobuf` fields.

When working with `protobuf`, it's quite common to have complex structures where a message field is contained within another message, which itself can be part of another message, and so on.
If these fields are accessed directly and some field in the call chain will not be initialized, it can result in application panic.

protosetter addresses this issue by suggesting use of setter methods for field updates.

**It is recommended to use [protogetter](https://github.com/ghostiam/protogetter) in conjunction with protosetter to ensure safe access to nested `protobuf` fields.**

## How does it work?
protosetter analyzes your Go code and helps detect direct `protobuf` field writes.\
The linter suggests using setters:
```go
m.GetFoo().GetBar().SetBaz(value)
```
instead of direct field writes:
```go
m.Foo.Bar.Baz = value
```

And you will then only need to perform a nil check before writing to the final field:
```go
if m.GetFoo().GetBar() != nil {
    m.GetFoo().GetBar().SetBaz(value)
}
```
instead of:
```go
if m.Foo != nil {
    if m.Foo.Bar != nil {
        if m.Foo.Bar != nil {
            m.Foo.Bar.Baz = value
        }
    }
}
```

which simplifies the code and makes it more reliable.

## Usage

Integration with [`golangci-lint`](https://github.com/golangci/golangci-lint) is planned.

## Standalone usage

### Installation

```bash
go install github.com/galexrt/protosetter/cmd/protosetter@latest
```

### Direct run

To run the linter:
```bash
protosetter ./...
```

Or to apply suggested fixes directly:
```bash
protosetter --fix ./...
```

## License

Based upon the code of [protogetter](https://github.com/ghostiam/protogetter) by [@ghostiam](https://github.com/ghostiam) which is licensed under the [MIT License](https://github.com/ghostiam/protogetter/blob/main/LICENSE).

Licensed under the [MIT License](LICENSE).
