# Microsoft Go examples

This repo has some basic sample apps that may be useful to exercise the [Microsoft build of Go](https://github.com/microsoft/go).
See also: <https://github.com/microsoft/go/blob/microsoft/main/eng/doc/fips/README.md>.

- [hello](hello) - Hello world
- [sum](sum) - Runs sha256 sums and responds with various details as a web server

## Build

### Local

1. Get binaries from <https://github.com/microsoft/go/blob/microsoft/main/eng/doc/Downloads.md>.
1. Go into an example dir.
1. Run `go run .`, where `go` is the path to the downloaded binary, `go/bin/go`.
1. Use `$env:GOEXPERIMENT = 'cngcrypto'` (pwsh) or `export GOEXPERIMENT=opensslcrypto` (bash) to use backends. (Among other ways.) Then `go run .` again to recompile and run.

### Main Dockerfile

Useful build + run for the main `Dockerfile`.
Uses a multi-stage build to keep the final image small, per standard practice.
Uses the Microsoft Go images so you don't need to download the Go toolset yourself.

```sh
docker build . -t sum && docker run -p 8080:8080 -it --rm sum
```

### Dockerfile.OneStage

This Dockerfile runs the app on the build machine.
Not good for deployment, but easier to use to prod at results because it includes Go, and a distro.

```sh
docker build . -t app -f Dockerfile.OneStage && docker run -p 8080:8080 -it --rm app
```

To poke around, override the entrypoint:

```sh
docker build . -t app -f Dockerfile.OneStage && docker run -p 8080:8080 -it --rm --entrypoint bash app
```

### Dockerfile.Windows

This builds on Windows.

```sh
docker build . -t app -f Dockerfile.Windows && docker run -p 8080:8080 -it --rm app
```
