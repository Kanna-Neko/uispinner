# uispinner
A Go library to render multiple spinners in terminal applications, support multi process

## feature
1. support multi process
2. chain method revise spinner
3. update spinner dynamic
4. every spinner has self interval

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