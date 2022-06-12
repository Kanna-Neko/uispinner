# uispinner
[![GoDoc](https://godoc.org/github.com/briandowns/spinner?status.svg)](https://pkg.go.dev/github.com/jaxleof/uispinner@v0.0.5) [![Go](https://github.com/jaxleof/uispinner/actions/workflows/go.yml/badge.svg)](https://github.com/jaxleof/uispinner/actions/workflows/go.yml)

A Go library to render multiple spinners in terminal applications, support multi process

## feature
1. support multi process
2. chain method revise spinner
3. update spinner dynamic
4. every spinner has self interval

## install
``` shell
go get github.com/jaxleof/uispinner
go get github.com/briandowns/spinner #this package supply many spinners
```

## usage
```go
cj := uispinner.New()
spinner1 := cj.AddSpinner(spinner.CharSets[34], 1*time.Millisecond).SetComplete("helloWorld").SetPrefix("abc").SetSuffix("ab")
spinner2 := cj.AddSpinner(spinner.CharSets[0], 100*time.Millisecond).SetComplete("good")
cj.Start()
time.Sleep(time.Second * 5)
spinner1.Done()
spinner2.Done()
cj.Stop()
```

## thanks
1. [uilive](https://github.com/gosuri/uilive). uilive is a go library for updating terminal output in realtime, which support whole uispinner.
2. [spinner](https://github.com/briandowns/spinner), which supply many interesting spinner.