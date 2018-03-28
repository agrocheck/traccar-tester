[![GoDoc](https://godoc.org/github.com/agrocheck/traccar-tester?status.svg)](https://godoc.org/github.com/agrocheck/traccar-tester)

# Traccar tester

TODO: description


## Build

```
go build -o traccar-tester cmd/traccar-tester/main.go
```


## Usage

```
nc -lk localhost 5031
```

```
./traccar-tester taip -host localhost -port 5031 -network tcp -count 5  -backoff 1 -id 1234 -event 14
```
