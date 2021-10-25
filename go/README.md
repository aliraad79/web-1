for running go individually install [Go](https://golang.org/doc/install) then comment line 31 of main.go and uncomment line 32 after that

```sh
    docker run -p 6379:6379 -d redis:6-alpine
    go run main.go
```