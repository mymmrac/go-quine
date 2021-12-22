# Go Quine

[Quine](https://en.wikipedia.org/wiki/Quine_(computing)) implementetion in Go.

> Inspired by: [Tsoding Daily: program that prints itself](https://youtu.be/GhYjEgRZjR8)

## Run

Generate new main:

```shell
go run main.go > new-main.go
```

Regenerate main:

```shell
go run new-main.go > main.go
```

## Check diff

```shell
diff -u main.go new-main.go
```



