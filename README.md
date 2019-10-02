# framework

## Unit Test

### Correr todos los tests
```go test ./... -coverprofile=cover.out```


### Correr todos los tests y ver detalles
```go test ./... -coverprofile=cover.out -v```

### Correr todos los tests y ver detalles en html
```go test ./... -coverprofile=cover.out && go tool cover -html=cover.out```