# security-compat

A shim that does exactly one thing:

provides symbol `_SecTrustCopyCertificateChain` (macOS 12+ API) on earlier macOS versions.

## Motivation

Enabling `go` toolchain and binaries, compiled by it, on `macOS Big Sur`.

While there are potentially more use cases, the above is the only one, tested by the author.

## Building

### With make (Xcode)

*NOTE*: Makefile sets `-DSECTRUST_COMPAT` to enable building the shim for macOS SDK 12+

```sh
make
```

### With nix

```sh
nix build
```

### Verify

The dylib must export `_SecTrustCopyCertificateChain` (and must not import it).

<details><summary>Click here for details</summary>

```sh
LIBSECCOMPAT=libsecurity_compat.dylib
# LIBSECCOMPAT=result/lib/libsecurity_compat.dylib

❯ nm -u "$LIBSECCOMPAT" | grep -o _SecTrustCopyCertificateChain

❯ nm -g "$LIBSECCOMPAT" | grep -o _SecTrustCopyCertificateChain
_SecTrustCopyCertificateChain

```

</details>

## Usage Examples

### Running go1.25 on macOS Big Sur

```sh
❯  DYLD_INSERT_LIBRARIES="$LIBSECCOMPAT" DYLD_FORCE_FLAT_NAMESPACE=1 go version
go version go1.25.4 darwin/arm64

```
