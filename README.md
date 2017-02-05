# Versioning

A library to make set/get/writing versioning information easier.

## Usage

1. Add in `main.go`:

```go
var version string

func init() {
    versioning.Set(version)
}

```

1. Where you want to get the version:

```go
versioning.Write(os.Stdout)
```

or

```go
version := versioning.Get()
```

1. Build your binary with this flag:

```bash
$ go build -ldflags "-X main.version=123-up-to-you"
```
