# Chafa Go - Conversion from C to Go

This directory contains the Go conversion of the Chafa library, a terminal graphics library that converts images to text-based representations.

## Status

🚧 **WORK IN PROGRESS** 🚧

This is an ongoing conversion of the Chafa C library to Go. The conversion is being done incrementally, starting with the core types and structures.

**Progress: Phase 2 Complete - Core Data Structures Converted**

### Statistics
- **Go Files**: 16 files
- **Lines of Code**: ~2,159 lines
- **Tests**: 24 tests, all passing
- **Original C Code**: ~40,000 lines
- **Completion**: ~5% (foundational structures complete)

### Completed

**Phase 1: Project Setup** ✅
- ✅ Go module structure (go.mod)
- ✅ Directory structure
- ✅ Build system

**Phase 2: Core Data Structures** ✅
- ✅ Common types and enums (PixelType, Align, Tuck, ColorExtractor, ColorSpace, DitherMode, etc.)
- ✅ Canvas configuration structure with all getters/setters
- ✅ Symbol map structure (partial implementation)
- ✅ Basic Canvas structure (stub)
- ✅ Color handling (internal/color.go)
  - Color structures and operations
  - RGB to DIN99d color space conversion
  - Color averaging and difference calculations
- ✅ Base64 encoding with streaming support (internal/base64.go)
- ✅ Math utilities (internal/math_util.go)
  - TuckAndAlign for image placement
  - Clamp, Min, Max, Square functions
  - RoundUpToMultipleOf
- ✅ Canvas geometry calculation (util.go)
  - CalcCanvasGeometry for optimal canvas sizing
  - Aspect ratio preservation
  - Font ratio compensation
- ✅ Version constants

### In Progress

- 🔄 Symbol map implementation
- 🔄 Terminal info structures

### Todo

**Core Components:**
- ⬜ Frame handling implementation
- ⬜ Image handling implementation
- ⬜ Placement implementation
- ⬜ Parser
- ⬜ Stream reader
- ⬜ Stream writer
- ⬜ Features detection
- ⬜ Remaining utilities

**Internal Components:**
- ⬜ Canvas printer
- ⬜ Pixel operations
- ⬜ Dithering algorithms
- ⬜ Indexed image support
- ⬜ Palette management
- ⬜ Color hash/table
- ⬜ Symbol definitions
- ⬜ Sixel canvas
- ⬜ Kitty canvas
- ⬜ iTerm2 canvas
- ⬜ Batch processing
- ⬜ Work cells and threading
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
  common.go         - Common types and enums ✅
  version.go        - Version constants ✅
  canvas.go         - Canvas rendering (stub) ✅
  canvas_config.go  - Canvas configuration ✅
  symbol_map.go     - Symbol mapping (partial) ✅
  term_info.go      - Terminal information (stub) ✅
  frame.go          - Frame handling (stub) ✅
  image.go          - Image handling (stub) ✅
  placement.go      - Placement (stub) ✅
  util.go           - Utilities ✅
  doc.go            - Package documentation ✅
  chafa_test.go     - Tests (9 tests) ✅
  
  internal/         - Internal implementation packages
    color.go        - Color handling ✅
    base64.go       - Base64 encoding ✅
    math_util.go    - Math utilities ✅
    internal_test.go - Tests (15 tests) ✅

cmd/chafa/          - CLI tool
  main.go           - Example application ✅
```

## Test Coverage

All 24 tests passing:
- ✅ Canvas configuration (creation, setters/getters, copying)
- ✅ Symbol map (creation, settings)
- ✅ Canvas creation
- ✅ Type constants (PixelType, SymbolTags, Optimizations)
- ✅ Canvas geometry calculation
- ✅ Color operations (pack/unpack, average, diff, accumulation)
- ✅ RGB to DIN99d conversion
- ✅ Base64 encoding (single and streaming)
- ✅ Math utilities (square, min/max, clamp, round up)
- ✅ Image placement calculations (TuckAndAlign)

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
go test ./chafa-go/...
```

## Running Example

```bash
go build -o /tmp/chafa-example ./cmd/chafa
/tmp/chafa-example
```

## Next Steps

1. **Phase 3**: Implement TermInfo and terminal database structures
2. **Symbol Definitions**: Convert symbol character mappings
3. **Pixel Operations**: Implement core pixel manipulation functions
4. **Canvas Rendering**: Complete the canvas drawing implementation
5. **Output Formats**: Implement Sixel, Kitty, and iTerm2 output

## License

LGPL 3.0 or later, same as the original Chafa library.

## Original C Library

The original C implementation can be found at: https://github.com/hpjansson/chafa
