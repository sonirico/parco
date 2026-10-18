# Changelog

## v0.14.0 - 2026-10-18

Everything in this release is additive and lives behind a `//go:build go1.27`
constraint. `go.mod` stays at Go 1.22; on older toolchains the new symbols do
not exist and the existing API is unchanged. Hand-written parsers and
compilers remain supported.

### Added

- `Model[T]` with the generic method `Field[U](Type[U], getter, setter)`.
  One method replaces the per-type builder methods and the
  `XField`/`XFieldGetter`/`XFieldSetter` constructors. A `Model[T]` is a
  `Type[T]`, so nested structs are plain fields. `Parse`, `ParseBytes` and
  `Compile` live on the model.
- `Registry[ID]`, a `Type[Entry[ID]]` that maps ids to models.
  `Register[U](id, *Model[U])` is typed at compile time. Nests as a field.
- `cmd/parco-gen`: a code generator that reads `parco` struct tags and writes
  the `Model` for every tagged struct (`go generate` friendly, stdlib only).
  Wire types are inferred from Go types; tags adjust byte order, header size,
  time location, or skip a field.
- `examples/codegen` with the generator input, its committed output and a
  roundtrip test; the generator's golden test regenerates it.
- CI matrix covers Go 1.26 and 1.27 so the constrained files are compiled and
  tested.

### Changed

- `OptionalType` implements `ByteLength`, so `Option(...)` is a `Type[U]` and
  can be passed to `Model.Field`.

### Measured

Compile benchmark, `Compiler` versus `Model` over the same struct, 10 runs on
go1.27.2: small 450.4 ns vs 447.5 ns, medium 2.546 us vs 2.548 us, large
25.23 us vs 25.24 us. Bytes, allocations and payload identical between the
two in every row.
