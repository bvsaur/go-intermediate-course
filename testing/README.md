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