# Testing in Go

## Run all defineed tests
```bash
go test
```

## Get test coverange in %
```bash
go test -cover
```

## Get a file with the test coverage of your program
```bash
go test -coverprofile=coverage.out
```

## Read coverage file on terminal
```bash
go tool cover -func=coverage.out
```

## Read coverage file on browser
```bash
go tool cover -html=coverage.out
```

## Run test with profiling
```bash
go test -cpuprofile=cpu.out
```

## Read profiling file
```bash
go tool pprof cpu.out
```

```bash
# List top time consuming processes
(pprof) top 

# Show average execution time of an especific function
(pprof) list [func_name] 

# Generate svg file to check visually the profiling
(pprof) web

# Generate PDF file to check visually the profiling
(pprof) pdf
```