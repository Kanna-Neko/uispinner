package main

import (
	"time"

	"github.com/jaxleof/uispinner"

	"github.com/briandowns/spinner"
)

func main() {
	cj := uispinner.New()
	spinner1 := cj.AddSpinner(spinner.CharSets[34], 1*time.Millisecond).SetComplete("helloWorld").SetPrefix("abc").SetSuffix("ab")
	spinner2 := cj.AddSpinner(spinner.CharSets[0], 100*time.Millisecond).SetComplete("good")
	cj.Start()
	time.Sleep(time.Second * 5)
	spinner1.Done()
	spinner2.Done()
	cj.Stop()
}
