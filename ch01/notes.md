

## Getting started notes:

### Env Vars
Be sure to set `GOPATH` for your src code and add go's bin folder to your `PATH`

```sh
export GOPATH=$(go env GO
```

```sh
export PATH=$PATH:$(go env GOPATH)/bin
```

### Main
For an executable program, it must be part of `package main` and have a `func() main`.


### loops
The only type of loop is a for loop. Can use this to mimic the other types:

#### while loop
```go
for x == false {
    // do something
}
```

#### infinite loop
```go
for {
    // do something forever
}
```

#### normal for loop:
```go
for i := 0; i < 10; i++ {
    // do something 10 times
}
```

#### iterate a collection:
```go
for index, value := range foo {
    // do something len(foo) times
}
```
