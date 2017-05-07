

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
