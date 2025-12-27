# security-compat

A shim that does exactly one thing:

provides symbol `_SecTrustCopyCertificateChain` (macOS 12+ API) on earlier macOS versions.

## Motivation

Enabling `go` toolchain and binaries, compiled by it, on `macOS Big Sur`.

While there are potentially more use cases, the above is the only one, tested by the author.

## Build

Make sure either Xcode or `nix` is available. Build using `make` or `nix build`.

## Usage Examples

### Running go1.25 on macOS Big Sur

```sh
DYLD_INSERT_LIBRARIES=libsecurity_compat.dylib DYLD_FORCE_FLAT_NAMESPACE=1 go build ./...
```

### Running a project built with go1.25 on macOS Big Sur

TODO
