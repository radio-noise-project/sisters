# Sisters 

## Setup
```bash
# Create empty .env file
touch .env
# download Go dependencies
go mod download
```
### Run Linter
```bash
golangci-lint run ./...
# or 
make ci
```
auto fix 
```
golangci-lint run ./... --fix 
```

