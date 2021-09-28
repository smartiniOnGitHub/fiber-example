# fiber-example

Minimal web application (webapp) made in Go with the Fiber framework


## Build

Latest Go (golang) stable, 1.16.x or later.
Go dependencies as specified in module descriptor.

Make (cmake) would be useful to simplify command lines, but it's optional.
Docker (latest) to build and run via Docker containers, optional.
UPX (latest) to compress/shrink executable, optional but reccomended.


## Setup, Build and run (in a DEV environment)

Set runtime properties in environment variables, 
or for example write them in a '.envrc' file (automatically loaded by 'direnv').

In the [docs](./docs/) folder there is other documentation.

To run it (from sources) use:
```bash
go run src/main.go
```

To build it:
```bash
# normal build (not optimized)
# go build main.go
go build -o ./build/fiber-example ./src/main.go
# or to build optimized (maybe use the './dist/' folder instead)
go build -o ./build/fiber-example -ldflags="-s -w" ./src/main.go
```
and then run with:
```bash
# ./main
./build/fiber-example
```

To compress/shrink executable, [UPX](https://upx.github.io/) is recommended.
As a sample, compress generated executable with:
```bash
upx ./build/fiber-example
```

To delete generated executable:
```bash
rm -f ./build/fiber-example*
```

Note that to simplify development, all shell commands have been included in a 'Makefile', 
so you can run them via related task name, like:
```bash
make # same of the info task
make run # same as run-dev here
make run-dev # build in a temporary file and run it directly
make test # build in a temporary file and run it directly
make build-and-run # build and run from a compiled version
make build-optimized-and-run # build and run from an optimized compiled version
# etc ...
```
see other tasks in the 'Makefile', like:
cross-compilation for Windows/Mac/Linux, build and run via Docker Containers, etc.


# Usage

Run the executable and provide options via (optional) flags:
```bash
./fiber-example
# ./fiber-example -help # to show help
```
as always, Windows executable has a '.exe' excension.
For other flags, see output from application help.


## Requirements

Nothing special installed is mandatory.
The 'direnv' utility could be useful.

To run the application in a container, an updated version of Docker is needed.


## Note

Nothing.


----
