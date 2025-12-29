# Chafa - Go Conversion

This repository contains an ongoing conversion of the Chafa library from C to Go.

## About Chafa

Chafa is a command-line utility and library that converts image data, including
animated GIFs, into graphics formats or ANSI/Unicode character art suitable
for display in a terminal. It has broad feature support, allowing it to be
used on devices ranging from historical teleprinters to modern terminal
emulators and everything in between.

The original C library can be found at: https://github.com/hpjansson/chafa

## Go Conversion Status

🚧 **WORK IN PROGRESS** 🚧

This is a large-scale conversion project converting approximately 40,000 lines
of C code to Go. The conversion is being done incrementally.

### What's Completed

- ✅ Basic project structure and Go module setup
- ✅ Common types and enums (PixelType, CanvasMode, ColorSpace, etc.)
- ✅ Canvas configuration API
- ✅ Symbol map structure
- ✅ Basic Canvas structure
- ✅ Tests for completed components

### Directory Structure

```
chafa-go/           - Go library package
  *.go              - Core library files
  README.md         - Detailed conversion progress
  internal/         - Internal implementation (TODO)

cmd/chafa/          - CLI tool (in progress)
  main.go           - Example/demo application

chafa/              - Original C library (preserved)
```

## Building

### Requirements

- Go 1.21 or later

### Build the Library

```bash
go build ./chafa-go
```

### Run Tests

```bash
go test ./chafa-go
```

### Build Example

```bash
go build -o chafa-go-example ./cmd/chafa
./chafa-go-example
```

## Conversion Approach

The conversion follows these principles:

1. **Literal Translation**: Be as literal as possible in the conversion while
   maintaining idiomatic Go code
2. **Preserve API Structure**: Maintain the existing API structure and
   functionality where it makes sense
3. **Go Idioms**: Use Go best practices (error returns instead of GError,
   channels instead of thread pools, etc.)
4. **No External Dependencies**: Avoid external dependencies where possible,
   use Go standard library

### GLib to Go Mappings

| C/GLib Type | Go Equivalent |
|-------------|---------------|
| `gint`, `guint` | `int`, `uint` |
| `gchar` | `byte` or `rune` |
| `gboolean` | `bool` |
| `GList`, `GSList` | `[]T` (slices) |
| `GHashTable` | `map[K]V` |
| `GString` | `strings.Builder` |
| `GError` | `error` interface |
| `g_malloc`/`g_free` | Go garbage collector |

### Threading to Concurrency

| C/GLib | Go |
|--------|-----|
| `pthread_t`, `GThread` | goroutines |
| `pthread_mutex_t`, `GMutex` | `sync.Mutex` |
| `GThreadPool` | Worker pool with goroutines |
| `GAsyncQueue` | `chan T` |

## Contributing

This conversion is part of a fork and not intended to be merged upstream.
The focus is on creating a clean, idiomatic Go implementation.

## License

LGPL 3.0 or later, same as the original Chafa library.

## Documentation

See `chafa-go/README.md` for detailed conversion progress and technical notes.

Run `go doc` to view API documentation:

```bash
go doc github.com/mmulet/chafa/chafa-go
```
