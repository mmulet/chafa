# Chafa Go - Conversion from C to Go

This directory contains the Go conversion of the Chafa library, a terminal graphics library that converts images to text-based representations.

## Status

🚧 **WORK IN PROGRESS** 🚧

This is an ongoing conversion of the Chafa C library to Go. The conversion is being done incrementally, starting with the core types and structures.

### Completed

- ✅ Basic project structure
- ✅ Common types and enums (PixelType, Align, Tuck, ColorExtractor, ColorSpace, DitherMode, etc.)
- ✅ Symbol map types and structure (partial)
- ✅ Canvas configuration structure
- ✅ Basic Canvas structure (stub)
- ✅ Version constants

### In Progress

- 🔄 Symbol map implementation
- 🔄 Canvas implementation
- 🔄 Terminal info structures

### Todo

**Core Components:**
- ⬜ Frame handling
- ⬜ Image handling
- ⬜ Placement
- ⬜ Parser
- ⬜ Stream reader
- ⬜ Stream writer
- ⬜ Features detection
- ⬜ Utilities

**Internal Components:**
- ⬜ Canvas printer
- ⬜ Color handling
- ⬜ Color hash/table
- ⬜ Palette management
- ⬜ Dithering algorithms
- ⬜ Indexed image support
- ⬜ Pixel operations
- ⬜ Symbol definitions
- ⬜ Sixel canvas
- ⬜ Kitty canvas
- ⬜ iTerm2 canvas
- ⬜ Base64 encoding
- ⬜ Batch processing
- ⬜ Work cells and threading
- ⬜ Math utilities
- ⬜ Noise generation
- ⬜ PCA
- ⬜ Passthrough encoder
- ⬜ String utilities

**SIMD Optimizations:**
- ⬜ Pure Go implementations to replace AVX2/SSE4.1/MMX/POPCNT
- ⬜ Consider Go assembly for performance-critical paths

**CLI Tool:**
- ⬜ Main CLI application
- ⬜ Image loaders (PNG, JPEG, GIF, etc.)
- ⬜ Integration with Go's image libraries

## Package Structure

```
chafa-go/           - Main package with public API
  common.go         - Common types and enums
  version.go        - Version constants
  canvas.go         - Canvas rendering
  canvas_config.go  - Canvas configuration
  symbol_map.go     - Symbol mapping
  term_info.go      - Terminal information (TODO)
  term_db.go        - Terminal database (TODO)
  frame.go          - Frame handling (TODO)
  image.go          - Image handling (TODO)
  placement.go      - Placement (TODO)
  parser.go         - Parsing (TODO)
  stream_reader.go  - Stream reading (TODO)
  stream_writer.go  - Stream writing (TODO)
  features.go       - Feature detection (TODO)
  util.go           - Utilities (TODO)
  
  internal/         - Internal implementation packages
    (to be created)

cmd/chafa/          - CLI tool
  main.go           - Main entry point (TODO)
```

## Conversion Notes

### GLib Replacements

The conversion replaces GLib types and functions with Go equivalents:

- `gint`, `guint`, etc. → `int`, `uint`
- `gchar` → `byte` or `rune`
- `gboolean` → `bool`
- `gpointer` → `unsafe.Pointer` or `interface{}`
- `GList`/`GSList` → `[]T` (slices)
- `GHashTable` → `map[K]V`
- `GString` → `strings.Builder`
- `GBytes`/`GByteArray` → `[]byte`
- `g_malloc`/`g_free` → Go's garbage collector
- `g_return_if_fail` → early return with if checks
- `GError` → `error` return values

### Threading

Threading is replaced with Go concurrency primitives:

- `pthread_t`/`GThread` → goroutines
- `pthread_mutex_t`/`GMutex` → `sync.Mutex`
- `pthread_cond_t`/`GCond` → `sync.Cond` or channels
- `GThreadPool` → Worker pool pattern
- `GAsyncQueue` → `chan T`
- `g_once`/`pthread_once` → `sync.Once`
- `GRWLock` → `sync.RWMutex`

### Memory Management

- No explicit allocation/deallocation needed
- Go's garbage collector handles memory
- Reference counting removed
- Consider `sync.Pool` for hot paths

## Building

```bash
go build ./chafa-go
```

## Testing

```bash
go test ./chafa-go
```

## License

LGPL 3.0 or later, same as the original Chafa library.

## Original C Library

The original C implementation can be found at: https://github.com/hpjansson/chafa
