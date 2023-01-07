# Go Modules


## Initialize a Go Module
```bash
go mod init github.com/[username]/[module-name]
```

## Install an external package
```bash
go get [package-url]
```

## Clean unused dependencies
```bash
go mod tidy
```

## Check cached dependecies installation path
```bash
go mod download -json
```