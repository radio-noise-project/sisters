# ⚡️Sisters 
**Sisters** is a client tool of **Last-Order**.  
This program provide a API client from Last-Order.

## Installation (For developers)
### Requirements
This project depends on below:  
- Docker 26.1.4
- Docker-compose 2.27.1

If you have not installed yet, please install these software.

### How to build and install
First, you download this git repository via `git clone`.
```
git clone git@github.com:radio-noise-project/sisters.git
```

This programs can be started with the `docker compose` command.  
Please execute below in `sisters` directory:  
```
docker compose up
```

### How to install the development environment
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

## License
