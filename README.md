# ra27-telemetry

## Get dependencies
`$ cd ra27-telemetry/`

`$ go get`

## Build
Build for x86.

`$ go build -o ra27-telemetry`

Build and static link

`$ go build -ldflags "-linkmode external -extldflags -static" -o ra27-telemetry`

## Run
Run

`$ ./ra27-telemetry run`
