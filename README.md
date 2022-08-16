# uispinner
![Coverage](https://img.shields.io/badge/Coverage-70.3%25-yellow)
[![GoDoc](https://godoc.org/github.com/briandowns/spinner?status.svg)](https://pkg.go.dev/github.com/jaxleof/uispinner@v0.0.5) [![Go](https://github.com/jaxleof/uispinner/actions/workflows/go.yml/badge.svg)](https://github.com/jaxleof/uispinner/actions/workflows/go.yml)

A Go library to render multiple spinners in terminal applications, support multi process and tree structure.

## feature
1. support multi process
2. tree structure supported
3. chain method revise spinner
4. update spinner dynamic
5. every spinner has self interval

## install
``` shell
go get github.com/jaxleof/uispinner
go get github.com/briandowns/spinner #this package supply many spinners
```

## usage
```go
cj := uispinner.New()
// Only multiples of 50*time.Millisecond are supported because io fresh is slow
spinner1 := cj.AddSpinner(spinner.CharSets[34], 50*time.Millisecond).SetComplete("helloWorld").SetPrefix("abc").SetSuffix("ab")
spinner2 := cj.AddSpinner(spinner.CharSets[0], 100*time.Millisecond).SetComplete("good")
cj.Start()
time.Sleep(time.Second * 5)
spinner1.Done()
spinner2.Done()
cj.Stop()
```

## snapshot
![snap.png](/snapshot.png)

## thanks
1. [uilive](https://github.com/gosuri/uilive). uilive is a go library for updating terminal output in realtime, which support whole uispinner.
2. [spinner](https://github.com/briandowns/spinner), which supply many interesting spinner.