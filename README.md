# Main Go Commands

## Running Code
```bash
go run .        # Run main package in current directory
go run main.go  # Run specific file
```

## Module Management
```bash
go mod init <module-name>  # Initialize a new module
go mod tidy               # Add missing and remove unused modules
go get <package>          # Add dependency
go mod download           # Download modules to local cache
```

## Building
```bash
go build          # Build executable in current directory
go build -o name  # Build with custom executable name
go install        # Build and install to $GOPATH/bin
```

## Testing
```bash
go test           # Run tests (files ending in _test.go, functions starting with Test)
go test -v        # Run tests with verbose output
go test ./...     # Run tests in current package and all sub-packages
go test -cover    # Run tests with coverage report
```

## Other Useful Commands
```bash
go fmt ./...      # Format all Go files
go vet            # Report suspicious constructs
go clean          # Remove build artifacts
go version        # Show Go version
go env            # Show Go environment variables
go list -f '{{.Target}}'  # Show install target path
``` 
