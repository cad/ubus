# ra27-telemetry

## Build
Build for x86.
`$ go build -o ra27-telemetry`

link static

`$ go build -ldflags "-linkmode external -extldflags -static" -o ra27-telemetry`

## Run
Run
`$ ./ra27-telemetry run`
